package templating

import (
	"bytes"
	"embed"
	"strings"
	"text/template"

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
