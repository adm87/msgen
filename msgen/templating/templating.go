package templating

import (
	"bytes"
	"embed"
	"strings"
	"text/template"
	"unicode"

	"github.com/adm87/msgen/models"
	"github.com/adm87/msgen/utils"
)

//go:embed templates
var templatesFS embed.FS

func loadTemplate(templatePath string) (string, error) {
	content, err := templatesFS.ReadFile(templatePath)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

var templatingFuncs = template.FuncMap{
	"buildRoutingTree": buildRoutingTree,
	"indent":           indent,
	"pascalCase":       pascalCase,
}

func RenderTemplate(name, templatePath string, data any) (string, error) {
	tmpl, err := loadTemplate(templatePath)

	if err != nil {
		return "", err
	}

	helpersTmpl, err := loadTemplate("templates/helpers.tpl")

	if err != nil {
		return "", err
	}

	t, err := template.New(name).Funcs(templatingFuncs).Parse(helpersTmpl + tmpl)
	t = t.Funcs(template.FuncMap{"include": includeTemplate(t)})

	if err != nil {
		return "", err
	}

	var outputBuffer bytes.Buffer

	if err = t.Execute(&outputBuffer, data); err != nil {
		return "", err
	}

	return outputBuffer.String(), nil
}

func RenderTemplateToFile(templatePath, outputPath string, data any) error {
	renderedContent, err := RenderTemplate(outputPath, templatePath, data)

	if err != nil {
		return err
	}

	return utils.WriteFile(outputPath, renderedContent)
}

func includeTemplate(t *template.Template) func(string, any) (string, error) {
	return func(n string, d any) (string, error) {
		buf := new(bytes.Buffer)
		tpl := t.Lookup(n)

		if tpl == nil {
			return "", nil
		}

		if err := tpl.Execute(buf, d); err != nil {
			return "", err
		}

		return buf.String(), nil
	}
}

func indent(text string, spaces int) string {
	prefix := strings.Repeat(" ", spaces)
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		if len(line) > 0 {
			lines[i] = prefix + line
		}
	}

	return strings.Join(lines, "\n")
}

func pascalCase(input string) string {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == '_' || r == '-' || unicode.IsSpace(r)
	})

	for i, w := range words {
		if len(w) == 0 {
			continue
		}

		runes := []rune(w)
		runes[0] = unicode.ToUpper(runes[0])
		for j := 1; j < len(runes); j++ {
			runes[j] = unicode.ToLower(runes[j])
		}
		words[i] = string(runes)
	}

	return strings.Join(words, "")
}

func buildRoutingTree(methods []models.SpecMethod) *models.SpecRoutingNode {
	root := &models.SpecRoutingNode{
		Children: make(map[string]*models.SpecRoutingNode),
	}

	for _, method := range methods {
		if method.Attributes == nil || method.Attributes.Router.Path == "" {
			continue
		}

		parts := strings.Split(method.Attributes.Router.Path[1:], "/")
		insertMethodIntoRoutingTree(root, parts, method)
	}

	return root
}

func insertMethodIntoRoutingTree(node *models.SpecRoutingNode, parts []string, method models.SpecMethod) {
	if len(parts) == 0 {
		node.Methods = append(node.Methods, method)
		return
	}

	part := parts[0]

	if _, exists := node.Children[part]; !exists {
		node.Children[part] = &models.SpecRoutingNode{
			Children: make(map[string]*models.SpecRoutingNode),
		}
	}

	insertMethodIntoRoutingTree(node.Children[part], parts[1:], method)
}
