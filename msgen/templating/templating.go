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

	t := template.New(name).Funcs(templatingFuncs)
	t = t.Funcs(template.FuncMap{
		"include": func(n string, d any) (string, error) {
			buf := new(bytes.Buffer)
			tpl := t.Lookup(n)

			if tpl == nil {
				return "", nil
			}

			if err := tpl.Execute(buf, d); err != nil {
				return "", err
			}

			return buf.String(), nil
		},
	})

	t, err = t.Parse(helpersTmpl + tmpl)

	if err != nil {
		return "", err
	}

	var outputBuffer bytes.Buffer

	if err = t.Execute(&outputBuffer, data); err != nil {
		return "", err
	}

	return outputBuffer.String(), nil
}
