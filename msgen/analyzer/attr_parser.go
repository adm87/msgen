package analyzer

import (
	"errors"
	"strings"

	"github.com/adm87/msgen/models"
)

var (
	ErrInvalidParamAttr    = errors.New("invalid parameter attribute, expected // @Param <name> <type> <in> <required>")
	ErrInvalidResponseAttr = errors.New("invalid response attribute, expected // @Success <code> <type> <dataType>")
	ErrInvalidRouterAttr   = errors.New("invalid router attribute, expected // @Router <path> [<method>]")
)

func parseParamAttr(attrStr string) (models.SpecMethodParamAttribute, error) {
	parts := strings.Split(attrStr, " ")

	if len(parts) != 4 {
		return models.SpecMethodParamAttribute{}, ErrInvalidParamAttr
	}

	fromPath := false
	fromQuery := false
	fromBody := false

	switch parts[2] {
	case "path":
		fromPath = true
	case "query":
		fromQuery = true
	case "body":
		fromBody = true
	}

	return models.SpecMethodParamAttribute{
		Name:      parts[0],
		Type:      parts[1],
		FromPath:  fromPath,
		FromQuery: fromQuery,
		FromBody:  fromBody,
		Required:  parts[3] == "true",
	}, nil
}

func parseResponseAttr(attrStr string) (models.SpecMethodResponseAttribute, error) {
	parts := strings.Split(attrStr, " ")

	if len(parts) != 3 {
		return models.SpecMethodResponseAttribute{}, ErrInvalidResponseAttr
	}

	return models.SpecMethodResponseAttribute{
		Code:     parts[0],
		Type:     parts[1],
		DataType: parts[2],
	}, nil
}

func parseRouterAttr(attrStr string) (models.SpecMethodRouterAttribute, error) {
	parts := strings.Split(attrStr, " ")

	if len(parts) != 2 {
		return models.SpecMethodRouterAttribute{}, ErrInvalidRouterAttr
	}

	return models.SpecMethodRouterAttribute{
		Path:   parts[0],
		Method: strings.ToUpper(strings.Trim(parts[1], "[]")),
	}, nil
}
