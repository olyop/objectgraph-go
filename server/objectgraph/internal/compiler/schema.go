package compiler

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/graphql-go/graphql"
	"github.com/vektah/gqlparser/ast"
)

func (c *compiler) compile() *graphql.SchemaConfig {
	_ = c.compileScalars()
	_ = c.compileTypes()

	return &graphql.SchemaConfig{
		Query: c.compileQuery(),
	}
}

func (c *compiler) compileScalars() map[string]*graphql.Object {
	var scalars = make(map[string]*graphql.Object)

	return scalars
}

func (c *compiler) compileTypes() map[string]*graphql.Object {
	var types = make(map[string]*graphql.Object)

	for typeName, typeDef := range c.ast.Types {
		if typeDef.BuiltIn {
			continue
		}

		if typeDef.Kind != ast.Object {
			continue
		}

		if typeName == "Query" || typeName == "Mutation" {
			continue
		}

		types[typeName] = graphql.NewObject(graphql.ObjectConfig{
			Name:   typeName,
			Fields: c.compileTypeFields(typeDef),
		})
	}

	return types
}

func (c *compiler) compileTypeFields(typeDef *ast.Definition) graphql.Fields {
	fields := graphql.Fields{}

	for _, fieldDef := range typeDef.Fields {
		fields[fieldDef.Name] = &graphql.Field{
			Name: fieldDef.Name,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		}
	}

	return fields
}

func (c *compiler) compileQuery() *graphql.Object {
	config := graphql.ObjectConfig{
		Name:   "Query",
		Fields: c.compileQueryFields(c.ast.Query),
	}

	return graphql.NewObject(config)
}

func (c *compiler) compileQueryFields(query *ast.Definition) graphql.Fields {
	fields := graphql.Fields{}

	for _, fieldDef := range query.Fields {
		fields[fieldDef.Name] = &graphql.Field{
			Name: fieldDef.Name,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		}
	}

	return fields
}

func (c *compiler) createResolver(groupKey string, retrieverKey string) func() error {
	retrievers := c.configuration.Retrievers[groupKey]

	// retrievers is a struct with retrievers as methods
	// use reflect to get the method
	funcValue := reflect.ValueOf(retrievers).MethodByName(retrieverKey)
	if !funcValue.IsValid() {
		panic(fmt.Sprintf("retriever %s not found", retrieverKey))
	}

	// check if the method has the right signature
	funcType := funcValue.Type()
	if funcType.NumIn() != 1 || funcType.NumOut() != 2 {
		panic(errors.New("retriever must have one input and two outputs"))
	}

	retriever := funcValue.Interface().(func(interface{}) (interface{}, error))

	return func() error {
		return nil
	}
}
