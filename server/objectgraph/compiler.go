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

	schemaConfig, err := sc.compile()
	if err != nil {
		return nil, err
	}

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
	objectsConfig objectConfigs
}

type objectConfigs map[string]*objectConfig
type objectConfig struct {
	objectKeyMap objectKeyMap
	objectType   reflect.Type
	retreivers   objectRetrievers
	graphqlType  *graphql.Object
}
type objectKeyMap map[string]string
type objectKey struct {
	typeName string
	value    string
}
type objectRetrievers map[string]reflect.Value

func (sc *schemaCompiler) compile() (graphql.SchemaConfig, error) {
	sc.createObjectKeyMaps()
	sc.createTypes()
	sc.createRetrievers()

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getProductByID": &graphql.Field{
				Type: sc.objectsConfig["Product"].graphqlType,
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
						return sc.mapData("Product", product)
					}

					input := RetrieverInput{
						PrimaryID: productID,
					}

					product, err = sc.execRetriever("Product", "ByID", input)
					if err != nil {
						return nil, err
					}

					sc.objectCache.Set("Product", productID.String(), product, time.Minute)

					return sc.mapData("Product", product)
				},
			},
			"getUsers": &graphql.Field{
				Type: graphql.NewList(sc.objectsConfig["User"].graphqlType),
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

	return schemaConfig, nil
}

const (
	directiveObjectKey         = "object"
	directiveObjectFieldArgKey = "field"
	directiveRetrieveKey       = "retrieve"
	directiveRetrieveKeyArgKey = "key"
)

func (sc *schemaCompiler) parseObjectsConfig() {
	objectsConfig := make(objectConfigs)

	for _, typeDef := range sc.ast.Types {
		if !isCustomType(typeDef) {
			continue
		}

		for _, fieldDef := range typeDef.Fields {
			objectDir := parseObjectDir(fieldDef)
			retreiveDir := parseRetrieveDir(fieldDef)

			if objectDir == nil && retreiveDir == nil {
				panic(fmt.Sprintf("directive not found for field %s", fieldDef.Name))
			}

			var typeName string
			if objectDir != nil {
				typeName = objectDir.key.typeName
			} else {
				typeName = retreiveDir.key.typeName
			}

			// initialize
			if objectsConfig[typeName] == nil {
				objectsConfig[typeName] = &objectConfig{}
			}

			// populate retriever
			if retreiveDir != nil {

			}
		}
	}

	sc.objectsConfig = objectsConfig
}

func parseObjectDir(fieldDef *ast.FieldDefinition) *objectDirective {
	objDirective := fieldDef.Directives.ForName(directiveObjectKey)
	if objDirective == nil {
		return nil
	}

	fieldArg := objDirective.Arguments.ForName(directiveObjectFieldArgKey)
	if fieldArg == nil {
		return nil
	}

	keyNotation := parseObjectKey(fieldArg.Value.Raw)

	return &objectDirective{
		key: keyNotation,
	}
}

type objectDirective struct {
	key *objectKey
}

func parseRetrieveDir(fieldDef *ast.FieldDefinition) *retrieveDirective {
	retrieveDir := fieldDef.Directives.ForName(directiveRetrieveKey)
	if retrieveDir == nil {
		return nil
	}

	retrieverKey := retrieveDir.Arguments.ForName(directiveRetrieveKeyArgKey)
	if retrieverKey == nil {
		return nil
	}

	key := parseObjectKey(retrieverKey.Value.Raw)

	return &retrieveDirective{
		key: key,
	}
}

type retrieveDirective struct {
	key *objectKey
}

func (sc *schemaCompiler) createObjectKeyMaps() {
	objectKeyMaps := make(objectKeyMaps)

	for _, typeDef := range sc.ast.Types {
		if !isCustomType(typeDef) {
			continue
		}

		structMap := make(map[string]string)
		for _, fieldDef := range typeDef.Fields {
			objDirective := fieldDef.Directives.ForName(directiveObjectKey)
			if objDirective == nil {
				continue
			}

			fieldArg := objDirective.Arguments.ForName(directiveObjectFieldArgKey)
			if fieldArg == nil {
				continue
			}

			keyNotation := parseObjectKey(fieldArg.Value.Raw)

			structMap[fieldDef.Name] = keyNotation.Value
		}

		objectKeyMaps[typeDef.Name] = structMap
	}

	sc.objectKeyMaps = objectKeyMaps
}

var builtInScalars = map[string]*graphql.Scalar{
	"String":  graphql.String,
	"Int":     graphql.Int,
	"Float":   graphql.Float,
	"Boolean": graphql.Boolean,
	"ID":      graphql.ID,
}

func (sc *schemaCompiler) createTypes() {
	// initialize custom types fields
	for _, typeDef := range sc.ast.Types {
		if !isCustomType(typeDef) {
			continue
		}

		fields := make(graphql.Fields)

		for _, fieldDef := range typeDef.Fields {
			scalar, isBuiltIn := builtInScalars[fieldDef.Type.NamedType]
			if !isBuiltIn {
				scalar = sc.config.Scalars[fieldDef.Type.NamedType]
			}

			if scalar == nil {
				panic(fmt.Sprintf("scalar %s not found", fieldDef.Type.NamedType))
			}

			fields[fieldDef.Name] = &graphql.Field{
				Type: scalar,
			}
		}

		sc.objectsConfig[typeDef.Name].graphqlType = graphql.NewObject(graphql.ObjectConfig{
			Name:   typeDef.Name,
			Fields: fields,
		})
	}

	// initialize custom types relations
	for _, typeDef := range sc.ast.Types {
		if !isCustomType(typeDef) {
			continue
		}

		for _, fieldDef := range typeDef.Fields {
			if sc.objectsConfig[typeDef.Name].graphqlType[fieldDef.Type.NamedType] == nil {
				continue
			}

			// do something with the field
			if fieldDef.Name != "brand" {
				continue
			}

			// graphqlTypes[fieldDef.Name]
		}

		graphqlTypes[typeDef.Name] = graphType
	}

	sc.types = graphqlTypes
}

func (sc *schemaCompiler) createRetrievers() {
	r := make(objectConfigs)

	for _, typeDef := range sc.ast.Types {
		if !isCustomObject(typeDef) {
			continue
		}

		for _, fieldDef := range typeDef.Fields {
			retrieveDir := fieldDef.Directives.ForName(directiveRetrieveKey)
			if retrieveDir == nil {
				continue
			}

			retrieverKey := retrieveDir.Arguments.ForName(directiveRetrieveKeyArgKey)
			if retrieverKey == nil {
				continue
			}

			key := parseObjectKey(retrieverKey.Value.Raw)

			objConfig := sc.config.Objects[key.TypeName]
			if objConfig == nil {
				panic(fmt.Sprintf("Object %s not found", key.TypeName))
			}

			// use reflect to get the method by name
			retrieverReflect := reflect.ValueOf(objConfig.Retrievers)
			retriever := retrieverReflect.MethodByName(key.Value)
			if !retriever.IsValid() {
				panic(fmt.Sprintf("Retriever %s not found", key.Value))
			}

			returnType := retriever.Type().Out(0)
			if returnType.Kind() == reflect.Ptr {
				returnType = returnType.Elem()
			}

			if r[key.TypeName] == nil {
				r[key.TypeName] = make(objectConfig)
			}

			r[key.TypeName][key.Value] = &retrieverConfig{
				function:   retriever,
				returnType: returnType,
			}
		}
	}

	sc.objectsConfig = r
}

func (sc *schemaCompiler) execRetriever(objectType string, funcName string, input RetrieverInput) (any, error) {
	retriever := sc.objectsConfig[objectType][funcName]
	args := []reflect.Value{reflect.ValueOf(input)}

	resultRef := retriever.function.Call(args)
	if len(resultRef) != 2 {
		return nil, fmt.Errorf("retriever %s must return 2 values", funcName)
	}

	println(retriever.returnType.String())

	valueReflect := resultRef[0]
	errReflect := resultRef[1]

	if !errReflect.IsNil() {
		return nil, errReflect.Interface().(error)
	}

	if dataIsStruct(valueReflect) {
		value := structToMap(dataGetValue(valueReflect))

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

			valueElem := structToMap(dataGetValue(valueReflect))

			value[i] = valueElem
		}

		return value, nil
	}

	return nil, fmt.Errorf("retriever %s must return a struct or a slice", funcName)
}

func (sc *schemaCompiler) mapData(objectType string, value any) (map[string]any, error) {
	m := make(map[string]any)

	valueReflect := reflect.ValueOf(value)
	isMap := valueReflect.Kind() == reflect.Map

	for fieldName, keyNotation := range sc.objectKeyMaps[objectType] {
		var fieldValue reflect.Value
		if isMap {
			// using reflect get the key value
			fieldValue = valueReflect.MapIndex(reflect.ValueOf(keyNotation))
		} else {
			fieldValue = valueReflect.FieldByName(fieldName)
		}

		if !fieldValue.IsValid() {
			return nil, fmt.Errorf("field %s not found", fieldName)
		}

		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()
		}

		if fieldValue.Kind() == reflect.Invalid {
			m[fieldName] = nil
		} else {
			m[fieldName] = fieldValue.Interface()
		}
	}

	return m, nil
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

func parseObjectKey(value string) *objectKey {
	parts := strings.Split(value, "/")

	if len(parts) != 2 {
		panic(fmt.Sprintf("Invalid key notation: %s", value))
	}

	return &objectKey{
		typeName: parts[0],
		value:    parts[1],
	}
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
		return value.Elem().Interface()
	}

	return value.Interface()
}

func dataGetSlice(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Ptr {
		return value.Elem()
	}

	return value
}

func structToMap(s any) map[string]any {
	m := make(map[string]any)

	sType := reflect.TypeOf(s)
	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		fieldValue := reflect.ValueOf(s).Field(i)

		m[field.Name] = fieldValue.Interface()
	}

	return m
}
