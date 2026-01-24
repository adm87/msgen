package templating

import (
	"bytes"
	"embed"
	"text/template"
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
	// Add custom template functions here if needed
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
