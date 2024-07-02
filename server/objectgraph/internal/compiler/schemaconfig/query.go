package schemaconfig

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/vektah/gqlparser/ast"
)

func (parser *schemaConfigParser) query() *graphql.Object {
	config := graphql.ObjectConfig{
		Name:   "Query",
		Fields: parser.queryFields(parser.ast.Query),
	}

	return graphql.NewObject(config)
}

func (parser *schemaConfigParser) queryFields(query *ast.Definition) graphql.Fields {
	fields := graphql.Fields{}

	for _, field := range query.Fields {
		if field.Position == nil {
			continue
		}

		if field.Name != "getProductByID" {
			continue
		}

		fields[field.Name] = parser.getProductByID(field)
	}

	return fields
}

func (parser *schemaConfigParser) getProductByID(field *ast.FieldDefinition) *graphql.Field {
	directives := parseDirectives(field.Directives)

	typeName := directives.RetrieverKey.ObjectType

	return &graphql.Field{
		Name: field.Name,
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			productID := p.Args["productID"].(string)

			product, exists, err := parser.objectcache.Get(typeName, productID, time.Second*10)
			if err != nil {
				return nil, err
			}

			if exists {
				return product, nil
			}

			return nil, nil
		},
	}
}
