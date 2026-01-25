package analyzer

import (
	"strings"

	"github.com/adm87/msgen/models"
)

func parseParamAttr(attrStr string) models.SpecMethodParamAttribute {
	parts := strings.Split(attrStr, " ")

	return models.SpecMethodParamAttribute{
		Name:     parts[0],
		Type:     parts[1],
		In:       parts[2],
		Required: parts[3] == "true",
	}
}

func parseResponseAttr(attrStr string) models.SpecMethodResponseAttribute {
	parts := strings.Split(attrStr, " ")

	return models.SpecMethodResponseAttribute{
		Code:     parts[0],
		Type:     parts[1],
		DataType: parts[2],
	}
}

func parseRouterAttr(attrStr string) models.SpecMethodRouterAttribute {
	parts := strings.Split(attrStr, " ")

	return models.SpecMethodRouterAttribute{
		Path:   parts[0],
		Method: strings.ToUpper(strings.Trim(parts[1], "[]")),
	}
}
