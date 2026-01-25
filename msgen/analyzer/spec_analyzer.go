package analyzer

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/adm87/msgen/models"
)

var (
	ErrInvalidTypeSpecCount = errors.New("invalid type spec count, expecting exactly one interface type")
	ErrInvalidSpecType      = errors.New("invalid spec type, expecting an interface type")
	ErrSpecMissingMethods   = errors.New("specification interface missing methods")
	ErrEmbeddedInterfaces   = errors.New("embedded interfaces are not supported")
	ErrMultipleMethodNames  = errors.New("method must have exactly one name")
)

// SpecAnalyzer implements the ast.Visitor interface for analyzing service specifications.
type SpecAnalyzer struct {
	Spec *models.Spec // Analyzed service information
	err  error        // Error encountered during analysis
}

func (sa *SpecAnalyzer) Visit(node ast.Node) ast.Visitor {
	if sa.err != nil {
		return nil
	}

	switch n := node.(type) {
	case *ast.GenDecl:

		switch n.Tok {
		case token.IMPORT:
			sa.err = parseImports(n, sa.Spec)

		case token.TYPE:
			sa.err = parseTypeSpec(n, sa.Spec)
		}
	}

	return sa
}

func parseImports(genDecl *ast.GenDecl, Spec *models.Spec) error {
	for _, spec := range genDecl.Specs {
		importSpec := spec.(*ast.ImportSpec)

		importPath := importSpec.Path.Value[1 : len(importSpec.Path.Value)-1]

		var alias string

		if importSpec.Name != nil {
			alias = importSpec.Name.Name
		}

		Spec.Imports = append(Spec.Imports, models.SpecImport{
			Path:  importPath,
			Alias: alias,
		})
	}

	return nil
}

func parseTypeSpec(genDecl *ast.GenDecl, Spec *models.Spec) error {
	if len(genDecl.Specs) != 1 {
		return ErrInvalidTypeSpecCount
	}

	typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec)

	if !ok {
		return ErrInvalidSpecType
	}

	interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)

	if !ok {
		return ErrInvalidSpecType
	}

	if len(interfaceType.Methods.List) == 0 {
		return ErrSpecMissingMethods
	}

	Spec.Name = typeSpec.Name.Name
	Spec.Attributes = parseSpecAttributes(collectAttributes(genDecl.Doc))

	for _, method := range interfaceType.Methods.List {
		funcType, ok := method.Type.(*ast.FuncType)

		if !ok {
			return ErrEmbeddedInterfaces
		}

		if len(method.Names) != 1 {
			return ErrMultipleMethodNames
		}

		methodAttrs, err := parseMethodAttributes(collectAttributes(method.Doc))

		if err != nil {
			return err
		}

		specMethod := models.SpecMethod{
			Name:       method.Names[0].Name,
			Attributes: methodAttrs,
			Parameters: parseFieldList(funcType.Params),
			Returns:    parseFieldList(funcType.Results),
		}

		Spec.Methods = append(Spec.Methods, specMethod)
	}

	return nil
}

func parseFieldList(fieldList *ast.FieldList) []models.SpecField {
	var params []models.SpecField

	if fieldList == nil {
		return params
	}

	for _, field := range fieldList.List {
		paramType := ""

		switch expr := field.Type.(type) {
		case *ast.Ident:
			paramType = expr.Name

		case *ast.SelectorExpr:
			paramType = formatSelectorType(expr, "")

		case *ast.StarExpr:
			if ident, ok := expr.X.(*ast.SelectorExpr); ok {
				paramType = formatSelectorType(ident, "*")
			}

		case *ast.ArrayType:
			if ident, ok := expr.Elt.(*ast.SelectorExpr); ok {
				paramType = formatSelectorType(ident, "[]")
			}
		}

		names := field.Names

		if len(names) == 0 {
			names = []*ast.Ident{{Name: ""}}
		}

		for _, name := range names {
			params = append(params, models.SpecField{
				Name: name.Name,
				Type: paramType,
			})
		}
	}

	return params
}

func parseSpecAttributes(attributes map[string][]string) *models.SpecAttributes {
	specAttrs := &models.SpecAttributes{}

	if summary, exists := attributes[models.SummaryAttr]; exists && len(summary) > 0 {
		specAttrs.Summary = summary[0]
	}

	if description, exists := attributes[models.DescriptionAttr]; exists && len(description) > 0 {
		specAttrs.Description = description[0]
	}

	return specAttrs
}

func parseMethodAttributes(attributes map[string][]string) (*models.SpecMethodAttributes, error) {
	methodAttrs := &models.SpecMethodAttributes{}

	if summary, exists := attributes[models.SummaryAttr]; exists && len(summary) > 0 {
		methodAttrs.Summary = summary[0]
	}

	if description, exists := attributes[models.DescriptionAttr]; exists && len(description) > 0 {
		methodAttrs.Description = description[0]
	}

	if version, exists := attributes[models.VersionAttr]; exists && len(version) > 0 {
		methodAttrs.Version = version[0]
	}

	if accepts, exists := attributes[models.AcceptsAttr]; exists && len(accepts) > 0 {
		methodAttrs.Accepts = accepts[0]
	}

	if returns, exists := attributes[models.ReturnsAttr]; exists && len(returns) > 0 {
		methodAttrs.Returns = returns[0]
	}

	if tags, exists := attributes[models.TagsAttr]; exists && len(tags) > 0 {
		methodAttrs.Tags = strings.Split(tags[0], " ")
	}

	if successAttrs, exists := attributes[models.SuccessAttr]; exists && len(successAttrs) > 0 {
		attr, err := parseResponseAttr(successAttrs[0])

		if err != nil {
			return nil, err
		}

		methodAttrs.Success = attr
	}

	if failureAttrs, exists := attributes[models.FailureAttr]; exists && len(failureAttrs) > 0 {
		for _, failureAttr := range failureAttrs {
			attr, err := parseResponseAttr(failureAttr)

			if err != nil {
				return nil, err
			}

			methodAttrs.Failures = append(methodAttrs.Failures, attr)
		}
	}

	if paramAttrs, exists := attributes[models.ParamAttr]; exists && len(paramAttrs) > 0 {
		for _, paramAttr := range paramAttrs {
			attr, err := parseParamAttr(paramAttr)

			if err != nil {
				return nil, err
			}

			methodAttrs.Parameters = append(methodAttrs.Parameters, attr)
		}
	}

	if routerAttrs, exists := attributes[models.RouterAttr]; exists && len(routerAttrs) > 0 {
		attr, err := parseRouterAttr(routerAttrs[0])

		if err != nil {
			return nil, err
		}

		methodAttrs.Router = attr
	}

	return methodAttrs, nil
}

func formatSelectorType(expr *ast.SelectorExpr, prefix string) string {
	if xIdent, ok := expr.X.(*ast.Ident); ok {
		return prefix + xIdent.Name + "." + expr.Sel.Name
	}
	return prefix + expr.Sel.Name
}

func collectAttributes(commentGroup *ast.CommentGroup) map[string][]string {
	attributes := make(map[string][]string)

	if commentGroup == nil {
		return attributes
	}

	for _, comment := range commentGroup.List {
		if len(comment.Text) > 2 && comment.Text[:2] == "//" {
			text := comment.Text[2:]

			if len(text) > 1 && text[0] == ' ' {
				text = text[1:]
			}

			if len(text) > 0 && text[0] == '@' {
				parts := strings.SplitN(text[1:], " ", 2)

				key := parts[0]
				value := ""

				if len(parts) > 1 {
					value = parts[1]
				}

				attributes[key] = append(attributes[key], value)
			}
		}
	}

	return attributes
}

// AnalyzeSpec analyzes the service specification file and returns Spec.
func AnalyzeSpec(specFile string) (*models.Spec, error) {
	scr, err := os.ReadFile(specFile)

	if err != nil {
		return nil, err
	}

	file, err := parser.ParseFile(token.NewFileSet(), specFile, scr, parser.ParseComments)

	if err != nil {
		return nil, err
	}

	visitor := &SpecAnalyzer{
		Spec: &models.Spec{
			Package: file.Name.Name,
			Imports: make([]models.SpecImport, 0),
			Methods: make([]models.SpecMethod, 0),
		},
	}

	ast.Walk(visitor, file)

	return visitor.Spec, visitor.err
}
