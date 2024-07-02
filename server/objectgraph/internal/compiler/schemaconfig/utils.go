package schemaconfig

import (
	"strings"
	"unicode"

	"github.com/vektah/gqlparser/ast"
)

func parseDirectives(directives ast.DirectiveList) *DirectivesConfig {
	return &DirectivesConfig{
		IsPrimary:    compileDirectiveIsPrimary(directives.ForName("primary")),
		DataField:    compileDirectiveDataField(directives.ForName("data")),
		RetrieverKey: compileDirectiveRetrieverKeys(directives.ForName("retrieve")),
	}
}

func compileDirectiveIsPrimary(primaryDirective *ast.Directive) bool {
	if primaryDirective == nil {
		return false
	} else {
		return true
	}
}

func compileDirectiveDataField(dataDirective *ast.Directive) *DirectivesObjectConfig {
	if dataDirective == nil {
		return nil
	}

	fieldArg := dataDirective.Arguments.ForName("field")
	if fieldArg == nil {
		return nil
	}

	return parseObjectReference(fieldArg.Value.Raw)
}

func compileDirectiveRetrieverKeys(retrieveDirective *ast.Directive) *DirectivesObjectConfig {
	if retrieveDirective == nil {
		return nil
	}

	keysArg := retrieveDirective.Arguments.ForName("key")
	if keysArg == nil {
		return nil
	}

	return parseObjectReference(keysArg.Value.Raw)
}

func parseObjectReference(input string) *DirectivesObjectConfig {
	split := strings.Split(input, "/")

	if len(split) != 2 {
		return nil
	}

	objectType := split[0]
	if !isAlphaNumeric(objectType) {
		return nil
	}

	objectField := split[1]
	if !isAlphaNumeric(objectField) {
		return nil
	}

	return &DirectivesObjectConfig{
		ObjectType:  objectType,
		ObjectField: objectField,
	}
}

func isAlphaNumeric(input string) bool {
	for _, r := range input {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return false
		}
	}

	return true
}

type DirectivesConfig struct {
	IsPrimary    bool
	DataField    *DirectivesObjectConfig
	RetrieverKey *DirectivesObjectConfig
}

type DirectivesObjectConfig struct {
	ObjectType  string
	ObjectField string
}
