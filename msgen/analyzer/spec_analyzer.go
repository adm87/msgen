package analyzer

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

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
	serviceInfo *models.ServiceInfo // Analyzed service information
	err         error               // Error encountered during analysis
}

func (sa *SpecAnalyzer) Visit(node ast.Node) ast.Visitor {
	if sa.err != nil {
		return nil
	}

	switch n := node.(type) {
	case *ast.GenDecl:

		switch n.Tok {
		case token.IMPORT:
			sa.err = parseImports(n, sa.serviceInfo)

		case token.TYPE:
			sa.err = parseTypeSpec(n, sa.serviceInfo)
		}
	}

	return sa
}

func parseImports(genDecl *ast.GenDecl, serviceInfo *models.ServiceInfo) error {
	for _, spec := range genDecl.Specs {
		importSpec := spec.(*ast.ImportSpec)

		importPath := importSpec.Path.Value[1 : len(importSpec.Path.Value)-1]

		var alias string

		if importSpec.Name != nil {
			alias = importSpec.Name.Name
		}

		serviceInfo.Imports = append(serviceInfo.Imports, models.SpecImport{
			Path:  importPath,
			Alias: alias,
		})
	}

	return nil
}

func parseTypeSpec(genDecl *ast.GenDecl, serviceInfo *models.ServiceInfo) error {
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

	for _, method := range interfaceType.Methods.List {
		funcType, ok := method.Type.(*ast.FuncType)
		if !ok {
			return ErrEmbeddedInterfaces
		}

		if len(method.Names) != 1 {
			return ErrMultipleMethodNames
		}

		specMethod := models.SpecMethod{
			Name:       method.Names[0].Name,
			Parameters: parseFieldList(funcType.Params),
			Returns:    parseFieldList(funcType.Results),
		}

		serviceInfo.Methods = append(serviceInfo.Methods, specMethod)
	}

	serviceInfo.Name = typeSpec.Name.Name
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

func formatSelectorType(expr *ast.SelectorExpr, prefix string) string {
	if xIdent, ok := expr.X.(*ast.Ident); ok {
		return prefix + xIdent.Name + "." + expr.Sel.Name
	}
	return prefix + expr.Sel.Name
}

// AnalyzeSpec analyzes the service specification file and returns ServiceInfo.
func AnalyzeSpec(specFile string) (*models.ServiceInfo, error) {
	scr, err := os.ReadFile(specFile)

	if err != nil {
		return nil, err
	}

	file, err := parser.ParseFile(token.NewFileSet(), specFile, scr, parser.ParseComments)

	if err != nil {
		return nil, err
	}

	visitor := &SpecAnalyzer{
		serviceInfo: &models.ServiceInfo{
			Package: file.Name.Name,
			Imports: make([]models.SpecImport, 0),
			Methods: make([]models.SpecMethod, 0),
		},
	}

	ast.Walk(visitor, file)

	return visitor.serviceInfo, visitor.err
}
