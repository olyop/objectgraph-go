package schemaconfig

import (
	"github.com/graphql-go/graphql"
	"github.com/olyop/objectgraph/objectgraph/internal/resolver"
	"github.com/vektah/gqlparser/ast"
)

func schemaCompileTypes(schema *ast.Schema) map[string]*graphql.Object {
	types := make(map[string]*graphql.Object)

	for _, t := range schema.Types {
		if t.Kind != ast.Object || t.BuiltIn {
			continue
		}

		config := graphql.ObjectConfig{
			Name:   t.Name,
			Fields: schemaCompileTypeFields(t.Fields),
		}

		types[t.Name] = graphql.NewObject(config)
	}

	return types
}

func schemaCompileTypeFields(fields ast.FieldList) graphql.Fields {
	graphqlFields := make(graphql.Fields)

	for _, field := range fields {
		directives := parseDirectives(field.Directives)

		f := &graphql.Field{
			Name: field.Name,
			Type: graphql.String,
			Resolve: resolver.Resolver(&resolver.ResolverOptions{
				IsPrimary: directives.IsPrimary,
				TypeName:  directives.DataField.ObjectType,
				DataKey:   directives.DataField.ObjectField,
			}),
		}

		graphqlFields[field.Name] = f
	}

	return graphqlFields
}
