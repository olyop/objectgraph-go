package objectgraph

import (
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/ast"
)

type compiler struct {
	schemaAst  *ast.Schema
	config     *Configuration
	retreivers retreivers

	schema *Schema
}

func compileSchema(schemaAst *ast.Schema, config *Configuration, retreivers retreivers) (*Schema, error) {
	c := &compiler{
		schemaAst:  schemaAst,
		config:     config,
		retreivers: retreivers,

		schema: &Schema{},
	}

	err := c.compile()
	if err != nil {
		return nil, err
	}

	return c.schema, nil
}

func (c *compiler) compile() error {
	for _, typeDef := range c.schemaAst.Types {
		switch typeDef.Kind {
		case ast.Scalar:
		case ast.Object:
			if typeDef.BuiltIn {
				continue
			}

			if typeDef.Name == "Query" {
				err := c.compileQuery(typeDef)
				if err != nil {
					return err
				}
			} else if typeDef.Name == "Mutation" {
			} else {
				err := c.compileType(typeDef)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (c *compiler) compileQuery(typeDef *ast.Definition) error {
	c.schema.Query = make(map[string]*SchemaField)

	for _, fieldDef := range typeDef.Fields {
		if strings.HasPrefix(fieldDef.Name, "__") {
			continue
		}

		ec, err := c.parseEngineConfig(fieldDef)
		if err != nil {
			return err
		}

		c.schema.Query[fieldDef.Name] = &SchemaField{
			SchemaType:   fieldDef.Type.Name(),
			FieldType:    determineFieldType(fieldDef.Type),
			EngineConfig: ec,
		}
	}

	return nil
}

func (c *compiler) compileType(typeDef *ast.Definition) error {
	if c.schema.Types == nil {
		c.schema.Types = make(map[string]map[string]*SchemaField)
	}

	if _, ok := c.schema.Types[typeDef.Name]; !ok {
		c.schema.Types[typeDef.Name] = make(map[string]*SchemaField)
	}

	for _, fieldDef := range typeDef.Fields {
		ec, err := c.parseEngineConfig(fieldDef)
		if err != nil {
			return err
		}

		c.schema.Types[typeDef.Name][fieldDef.Name] = &SchemaField{
			SchemaType:   fieldDef.Type.Name(),
			FieldType:    determineFieldType(fieldDef.Type),
			EngineConfig: ec,
		}
	}

	return nil
}

func (c *compiler) parseEngineConfig(fieldDef *ast.FieldDefinition) (*SchemaFieldEngine, error) {
	retrieve, err := c.parseEngineConfigRetrieve(fieldDef)
	if err != nil {
		return nil, err
	}

	object, err := c.parseEngineConfigObject(fieldDef)
	if err != nil {
		return nil, err
	}

	return &SchemaFieldEngine{
		Retriever: retrieve,
		Object:    object,
	}, nil
}

const (
	retrieveDirectiveName    = "retrieve"
	retrieveDirectiveArgKey  = "key"
	retrieveDirectiveArgArgs = "args"
)

func (c *compiler) parseEngineConfigRetrieve(fieldDef *ast.FieldDefinition) (*SchemaFieldEngineRetrieve, error) {
	retrieveDirective := fieldDef.Directives.ForName(retrieveDirectiveName)
	if retrieveDirective == nil {
		return nil, nil
	}

	keyArg := retrieveDirective.Arguments.ForName(retrieveDirectiveArgKey)
	if keyArg == nil {
		return nil, fmt.Errorf("key argument not found in retrieve directive")
	}
	objectKey, err := parseObjectKey(keyArg.Value.Raw)
	if err != nil {
		return nil, err
	}
	if _, ok := c.config.Objects[objectKey.TypeName]; !ok {
		return nil, fmt.Errorf("object %s not found in configuration", objectKey.TypeName)
	}
	if _, ok := c.retreivers[objectKey.TypeName].retreivers[objectKey.FieldKey]; !ok {
		return nil, fmt.Errorf("retriever %s not found in configuration", objectKey.FieldKey)
	}

	var args map[string]string
	argsArg := retrieveDirective.Arguments.ForName(retrieveDirectiveArgArgs)
	if argsArg != nil {
		args = make(map[string]string)
		for _, engineArg := range argsArg.Value.Children {
			argKey, argValue := parseEngineArg(engineArg.Value.Raw)
			args[argKey] = argValue
		}
	}

	return &SchemaFieldEngineRetrieve{
		Key:  objectKey,
		Args: args,
	}, nil
}

const (
	objectDirectiveName     = "object"
	objectDirectiveArgField = "key"
)

func (c *compiler) parseEngineConfigObject(fieldDef *ast.FieldDefinition) (*SchemaFieldEngineObject, error) {
	objectDirective := fieldDef.Directives.ForName(objectDirectiveName)
	if objectDirective == nil {
		return nil, nil
	}

	keyArg := objectDirective.Arguments.ForName(objectDirectiveArgField)
	if keyArg == nil {
		return nil, fmt.Errorf("key argument not found in object directive")
	}
	objectKey, err := parseObjectKey(keyArg.Value.Raw)
	if err != nil {
		return nil, err
	}
	if _, ok := c.config.Objects[objectKey.TypeName]; !ok {
		return nil, fmt.Errorf("object %s not found in configuration", objectKey.TypeName)
	}

	return &SchemaFieldEngineObject{
		Key: objectKey,
	}, nil
}

func parseObjectKey(keyValue string) (*SchemaObjectKey, error) {
	parts := strings.Split(keyValue, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid object key format: %s", keyValue)
	}

	return &SchemaObjectKey{
		TypeName: parts[0],
		FieldKey: parts[1],
	}, nil
}

// primaryID=$productID
func parseEngineArg(argValue string) (string, string) {
	parts := strings.Split(argValue, "=")
	if len(parts) != 2 {
		return "", ""
	}

	return parts[0], parts[1]
}

func determineFieldType(typeDef *ast.Type) *SchemaFieldType {
	var isArray bool
	var isArrayNonNull bool

	if typeDef.NamedType == "" {
		isArray = true
		isArrayNonNull = typeDef.Elem.NonNull
	}

	return &SchemaFieldType{
		NonNull:        typeDef.NonNull,
		IsArray:        isArray,
		IsArrayNonNull: isArrayNonNull,
	}
}
