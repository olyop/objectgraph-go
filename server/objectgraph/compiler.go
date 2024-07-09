package objectgraph

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/olyop/objectgraph/objectgraph/internal/objectcache"
	"github.com/vektah/gqlparser/ast"
)

func compileSchema(config *Configuration, ast *ast.Schema, objectcache *objectcache.ObjectCache) (*graphql.Schema, error) {
	sc := newSchemaCompiler(config, ast, objectcache)

	schemaConfig := sc.compile()

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}

func newSchemaCompiler(config *Configuration, ast *ast.Schema, objectcache *objectcache.ObjectCache) *schemaCompiler {
	sc := &schemaCompiler{
		config:      config,
		ast:         ast,
		objectCache: objectcache,
	}

	return sc
}

type schemaCompiler struct {
	config      *Configuration
	ast         *ast.Schema
	objectCache *objectcache.ObjectCache

	// internal
	structMappers structMappers
	types         map[string]graphql.Type
	retreivers    retreivers
}

type structMappers map[string]map[string]string
type retreivers map[string]map[string]reflect.Value

func (sc *schemaCompiler) compile() graphql.SchemaConfig {
	sc.createStructMappers()
	sc.createTypes()
	sc.createRetrievers()

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getProductByID": &graphql.Field{
				Type: sc.types["Product"],
				Args: graphql.FieldConfigArgument{
					"productID": &graphql.ArgumentConfig{
						Type: sc.config.Scalars["UUID"],
					},
				},
				Resolve: func(p graphql.ResolveParams) (any, error) {
					productID := p.Args["productID"].(uuid.UUID)

					product, exists, err := sc.objectCache.Get("Product", productID.String(), time.Minute)
					if err != nil {
						return nil, err
					}

					if exists {
						return product, nil
					}

					input := RetrieverInput{
						PrimaryID: productID,
					}

					product, err = sc.execRetriever("Product", "ByID", input)
					if err != nil {
						return nil, err
					}

					// sc.objectCache.Set("Product", productID.String(), product, time.Minute)

					return product, nil
				},
			},
			"getUsers": &graphql.Field{
				Type: graphql.NewList(sc.types["User"]),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					users, exists, err := sc.objectCache.Get("Query", "Top1000", time.Minute)
					if err != nil {
						return nil, err
					}

					if exists {
						return users, nil
					}

					users, err = sc.execRetriever("Query", "Top1000", RetrieverInput{})
					if err != nil {
						return nil, err
					}

					sc.objectCache.Set("Query", "Top1000", users, time.Minute)

					return users, nil
				},
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"clearCache": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					err := sc.objectCache.Clear()
					if err != nil {
						return false, err
					}

					return true, nil
				},
			},
		},
	})

	schemaConfig := graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	}

	return schemaConfig
}

const (
	directiveObjectKey         = "object"
	directiveObjectFieldArgKey = "field"
	directiveRetrieveKey       = "retrieve"
	directiveRetrieveKeyArgKey = "key"
)

func (sc *schemaCompiler) createStructMappers() {
	m := make(structMappers)

	for _, objDef := range sc.ast.Types {
		if !isCustomType(objDef) {
			continue
		}

		structMap := make(map[string]string)
		for _, fieldDef := range objDef.Fields {
			objDirective := fieldDef.Directives.ForName(directiveObjectKey)
			if objDirective == nil {
				continue
			}

			fieldArg := objDirective.Arguments.ForName(directiveObjectFieldArgKey)
			if fieldArg == nil {
				continue
			}

			keyNotation := parseKeyNotation(fieldArg.Value.Raw)

			structMap[fieldDef.Name] = keyNotation.Value
		}

		m[objDef.Name] = structMap
	}

	sc.structMappers = m
}

func (sc *schemaCompiler) createTypes() {
	graphqlTypes := make(map[string]graphql.Type)

	builtInScalars := map[string]*graphql.Scalar{
		"String":  graphql.String,
		"Int":     graphql.Int,
		"Float":   graphql.Float,
		"Boolean": graphql.Boolean,
		"ID":      graphql.ID,
	}

	// initialize custom types fields
	for _, typeDef := range sc.ast.Types {
		if !isCustomType(typeDef) {
			continue
		}

		fields := make(graphql.Fields)

		for _, fieldDef := range typeDef.Fields {
			scalar := builtInScalars[fieldDef.Type.NamedType]

			// check if the field is a custom scalar
			if scalar == nil {
				scalar = sc.config.Scalars[fieldDef.Type.NamedType]
			}

			if scalar == nil {
				continue
			}

			fields[fieldDef.Name] = &graphql.Field{
				Type: scalar,
			}
		}

		graphqlTypes[typeDef.Name] = graphql.NewObject(graphql.ObjectConfig{
			Name:   typeDef.Name,
			Fields: fields,
		})
	}

	// initialize custom types relations
	for _, typeDef := range sc.ast.Types {
		if !isCustomType(typeDef) {
			continue
		}

		graphType := graphqlTypes[typeDef.Name]

		for _, fieldDef := range typeDef.Fields {
			if graphqlTypes[fieldDef.Type.NamedType] == nil {
				continue
			}

			// do something with the field
		}

		graphqlTypes[typeDef.Name] = graphType
	}

	sc.types = graphqlTypes
}

func (sc *schemaCompiler) createRetrievers() {
	r := make(retreivers)

	for _, objDef := range sc.ast.Types {
		if !isCustomObject(objDef) {
			continue
		}

		typeRetrievers := make(map[string]reflect.Value)

		for _, fieldDef := range objDef.Fields {
			retrieveDir := fieldDef.Directives.ForName(directiveRetrieveKey)
			if retrieveDir == nil {
				continue
			}

			retrieverKey := retrieveDir.Arguments.ForName(directiveRetrieveKeyArgKey)
			if retrieverKey == nil {
				continue
			}

			key := parseKeyNotation(retrieverKey.Value.Raw)

			objConfig := sc.config.Objects[key.TypeName]
			if objConfig == nil {
				panic(fmt.Sprintf("Object %s not found", key.TypeName))
			}

			// use reflect to get the method by name
			retriever := reflect.ValueOf(objConfig.Retrievers).MethodByName(key.Value)
			if !retriever.IsValid() {
				panic(fmt.Sprintf("Retriever %s not found", key.Value))
			}

			typeRetrievers[key.Value] = retriever
		}

		r[objDef.Name] = typeRetrievers
	}

	sc.retreivers = r
}

func (sc *schemaCompiler) execRetriever(objectType string, funcName string, input RetrieverInput) (any, error) {
	retriever := sc.retreivers[objectType][funcName]
	if !retriever.IsValid() {
		return nil, fmt.Errorf("retriever %s not found", funcName)
	}

	args := []reflect.Value{reflect.ValueOf(input)}

	resultRef := retriever.Elem().Call(args)
	if len(resultRef) != 2 {
		return nil, fmt.Errorf("retriever %s must return 2 values", funcName)
	}

	valueReflect := resultRef[0]
	errReflect := resultRef[1]

	if !errReflect.IsNil() {
		return nil, errReflect.Interface().(error)
	}

	if dataIsStruct(valueReflect) {
		value, err := sc.mapData(objectType, dataGetValue(valueReflect))
		if err != nil {
			return nil, err
		}

		return value, nil
	}

	if dataIsSlice(valueReflect) {
		valueReflect = dataGetSlice(valueReflect)

		value := make([]map[string]any, valueReflect.Len())

		for i := 0; i < valueReflect.Len(); i++ {
			valueReflect = valueReflect.Index(i)

			if !dataIsStruct(valueReflect) {
				return nil, fmt.Errorf("retriever %s must return a struct or a pointer to a struct", funcName)
			}

			valueTmp, err := sc.mapData(objectType, dataGetValue(valueReflect))
			if err != nil {
				return nil, err
			}

			value[i] = valueTmp
		}

		return value, nil
	}

	return nil, fmt.Errorf("retriever %s must return a struct or a slice", funcName)
}

func isCustomType(def *ast.Definition) bool {
	if !isCustomObject(def) {
		return false
	}

	if def.Name == "Query" || def.Name == "Mutation" {
		return false
	}

	return true
}

func isCustomObject(def *ast.Definition) bool {
	if def.Kind != ast.Object {
		return false
	}

	if def.BuiltIn {
		return false
	}

	return true
}

func parseKeyNotation(value string) KeyNotation {
	// split by /
	parts := strings.Split(value, "/")

	if len(parts) != 2 {
		panic(fmt.Sprintf("Invalid key notation: %s", value))
	}

	return KeyNotation{
		TypeName: parts[0],
		Value:    parts[1],
	}
}

type KeyNotation struct {
	TypeName string
	Value    string
}

// check if is a struct or a pointer to a struct
func dataIsStruct(value reflect.Value) bool {
	if value.Kind() == reflect.Struct {
		return true
	}

	if value.Kind() == reflect.Pointer && value.Elem().Kind() == reflect.Struct {
		return true
	}

	return false
}

// check if is a slice or a pointer to a slice
func dataIsSlice(value reflect.Value) bool {
	if value.Kind() == reflect.Slice {
		return true
	}

	if value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Slice {
		return true
	}

	return false
}

func dataGetValue(value reflect.Value) any {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value.Interface()
}

func dataGetSlice(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}

func (sc *schemaCompiler) mapData(typeName string, value any) (map[string]any, error) {
	var m map[string]any

	valueReflect := reflect.ValueOf(value)

	for fieldName, keyNotation := range sc.structMappers[typeName] {
		fieldValue := valueReflect.FieldByName(keyNotation)

		if !fieldValue.IsValid() {
			return nil, fmt.Errorf("field %s not found", fieldName)
		}

		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()
		}

		m[fieldName] = fieldValue.Interface()
	}

	return m, nil
}
