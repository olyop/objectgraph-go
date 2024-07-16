package objectgraph

import (
	"fmt"
	"reflect"
	"time"

	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/olyop/objectgraph/database"
	"github.com/olyop/objectgraph/objectgraph/internal/objectcache"
)

func compileSchemaConfig(config *Configuration, objectcache *objectcache.ObjectCache) (*graphql.SchemaConfig, error) {
	compiler := engineCompiler{
		config:      config,
		objectcache: objectcache,
	}

	return compiler.createSchemaConfig()
}

type engineCompiler struct {
	config      *Configuration
	objectcache *objectcache.ObjectCache

	retrieversConfig retreiversConfig
}

func (engine *engineCompiler) createSchemaConfig() (*graphql.SchemaConfig, error) {
	err := engine.createRetrieversConfig()
	if err != nil {
		return nil, err
	}

	brandType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Brand",
		Fields: graphql.Fields{
			"brandID": &graphql.Field{
				Type: graphql.NewNonNull(engine.config.Scalars["UUID"]),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return objectField(p.Source, "Brand", "BrandID")
				},
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(builtInScalars["String"]),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return objectField(p.Source, "Brand", "Name")
				},
			},
			"updatedAt": &graphql.Field{
				Type: engine.config.Scalars["Timestamp"],
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return objectField(p.Source, "Brand", "UpdatedAt")
				},
			},
			"createdAt": &graphql.Field{
				Type: graphql.NewNonNull(engine.config.Scalars["Timestamp"]),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return objectField(p.Source, "Brand", "CreatedAt")
				},
			},
		},
	})

	categoryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"categoryID": &graphql.Field{
				Type: graphql.NewNonNull(engine.config.Scalars["UUID"]),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return objectField(p.Source, "Category", "CategoryID")
				},
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(builtInScalars["String"]),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return objectField(p.Source, "Category", "Name")
				},
			},
			"updatedAt": &graphql.Field{
				Type: engine.config.Scalars["Timestamp"],
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return objectField(p.Source, "Category", "UpdatedAt")
				},
			},
			"createdAt": &graphql.Field{
				Type: graphql.NewNonNull(engine.config.Scalars["Timestamp"]),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return objectField(p.Source, "Category", "CreatedAt")
				},
			},
		},
	})

	productType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"productID": &graphql.Field{
				Type: graphql.NewNonNull(engine.config.Scalars["UUID"]),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return objectField(p.Source, "Product", "ProductID")
				},
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(builtInScalars["String"]),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return objectField(p.Source, "Product", "Name")
				},
			},
			"updatedAt": &graphql.Field{
				Type: engine.config.Scalars["Timestamp"],
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return objectField(p.Source, "Product", "UpdatedAt")
				},
			},
			"createdAt": &graphql.Field{
				Type: graphql.NewNonNull(engine.config.Scalars["Timestamp"]),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return objectField(p.Source, "Product", "CreatedAt")
				},
			},
			"brand": &graphql.Field{
				Type: graphql.NewNonNull(brandType),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					brandIDValue, err := objectField(p.Source, "Product", "BrandID")
					if err != nil {
						return nil, err
					}

					brandID := brandIDValue.(uuid.UUID)

					object, exists, err := engine.objectcache.Get("Brand", engine.retrieversConfig["Brand"].typ, brandID.String(), time.Minute)
					if err != nil {
						return nil, err
					}
					if exists {
						return resolveForType("Brand", object), nil
					}

					input := RetrieverInput{
						PrimaryID: brandID,
					}
					object, err = engine.execRetriever("Brand", "ByID", input)
					if err != nil {
						return nil, err
					}

					engine.objectcache.Set("Brand", brandID.String(), object, time.Minute)

					return resolveForType("Brand", object), nil
				},
			},
			// "categories": &graphql.Field{
			// 	Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(categoryType))),
			// 	Resolve: func(p graphql.ResolveParams) (any, error) {
			// 		productIDValue, err := objectField(p.Source, "Product", "ProductID")
			// 		if err != nil {
			// 			return nil, err
			// 		}

			// 		productID := productIDValue.(uuid.UUID)

			// 		input := RetrieverInput{
			// 			PrimaryID: productID,
			// 		}
			// 		object, err := engine.execRetriever("Category", "AllByProductID", input)
			// 		if err != nil {
			// 			return nil, err
			// 		}

			// 		engine.objectcache.Set("Category", productID.String(), object, time.Minute)

			// 		return resolveForType("Category", object), nil
			// 	},
			// },
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getBrandByID": &graphql.Field{
				Type: graphql.NewNonNull(brandType),
				Args: graphql.FieldConfigArgument{
					"brandID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(engine.config.Scalars["UUID"]),
					},
				},
				Resolve: func(p graphql.ResolveParams) (any, error) {
					brandID := p.Args["brandID"].(uuid.UUID)

					object, exists, err := engine.objectcache.Get("Brand", engine.retrieversConfig["Brand"].typ, brandID.String(), time.Minute)
					if err != nil {
						return nil, err
					}
					if exists {
						return resolveForType("Brand", object), nil
					}

					result, err := database.SelectBrandByID(brandID)
					if err != nil {
						return nil, err
					}
					object = structs.Map(result)

					engine.objectcache.Set("Brand", brandID.String(), object, time.Minute)

					return resolveForType("Brand", object), nil
				},
			},
			"getProductByID": &graphql.Field{
				Type: graphql.NewNonNull(productType),
				Args: graphql.FieldConfigArgument{
					"productID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(engine.config.Scalars["UUID"]),
					},
				},
				Resolve: func(p graphql.ResolveParams) (any, error) {
					productID := p.Args["productID"].(uuid.UUID)

					object, exists, err := engine.objectcache.Get("Product", engine.retrieversConfig["Product"].typ, productID.String(), time.Minute)
					if err != nil {
						return nil, err
					}
					if exists {
						return resolveForType("Product", object), nil
					}

					result, err := database.SelectProductByID(productID)
					if err != nil {
						return nil, err
					}
					object = structs.Map(result)

					engine.objectcache.Set("Product", productID.String(), object, time.Minute)

					return resolveForType("Product", object), nil
				},
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"clearCache": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Boolean),
				Resolve: func(p graphql.ResolveParams) (any, error) {
					engine.objectcache.Clear()
					return true, nil
				},
			},
		},
	})

	schemaConfig := graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	}

	return &schemaConfig, nil
}

// createRetrieversConfig creates a map of retrieversObjectConfig from the configuration
// it checks that all methods on the retrievers object have the correct signature
// and that they all return the same type
func (engine *engineCompiler) createRetrieversConfig() error {
	rc := make(retreiversConfig)

	for typeName, objectConfig := range engine.config.Objects {
		var typ reflect.Type
		retrievers := make(retrievers)

		// using reflect, loop over all methods on the retrievers object
		retrieversValue := reflect.ValueOf(objectConfig.Retrievers)
		retrieversTyp := retrieversValue.Type()
		for i := 0; i < retrieversValue.NumMethod(); i++ {
			method := retrieversTyp.Method(i)
			methodVal := retrieversValue.Method(i)
			methodTyp := methodVal.Type()

			retrieverTyp, err := parseRetriever(methodTyp)
			if err != nil {
				return fmt.Errorf("%s: %s", typeName, err)
			}

			if typ == nil {
				typ = retrieverTyp
			} else if typ != retrieverTyp {
				return fmt.Errorf("%s: all methods should return the same type, got %s and %s", typeName, typ.String(), retrieverTyp.String())
			}

			retrievers[method.Name] = methodVal
		}

		rc[typeName] = &retrieversObjectConfig{
			primaryKey: objectConfig.PrimaryKey,
			typ:        typ,
			retreivers: retrievers,
		}
	}

	engine.retrieversConfig = rc

	return nil
}

func (engine *engineCompiler) execRetriever(objectType string, funcName string, input RetrieverInput) (any, error) {
	retriever := engine.retrieversConfig[objectType].retreivers[funcName]
	args := []reflect.Value{reflect.ValueOf(input)}

	resultValue := retriever.Call(args)

	// check error
	if len(resultValue) == 2 && !resultValue[1].IsNil() {
		return nil, resultValue[1].Interface().(error)
	}

	// check if the result is a slice
	if resultValue[0].Kind() == reflect.Slice {
		result := make([]map[string]any, resultValue[0].Len())
		for i := 0; i < resultValue[0].Len(); i++ {
			result[i] = structs.Map(resultValue[0].Index(i).Interface())
		}
		return result, nil
	} else {
		return structs.Map(resultValue[0].Interface()), nil
	}
}

func resolveForType(typeName string, value map[string]any) map[string]map[string]any {
	m := make(map[string]map[string]any)
	m[typeName] = value
	return m
}

func objectField(source any, typeName string, field string) (any, error) {
	objects, ok := source.(map[string]map[string]any)
	if !ok {
		return nil, fmt.Errorf("%s/%s: source is not a map[string]map[string]any", typeName, field)
	}

	// get the object from the source
	object, ok := objects[typeName]
	if !ok {
		return nil, fmt.Errorf("%s/%s: source does not contain %s", typeName, field, typeName)
	}

	// get the field from the object
	value, ok := object[field]
	if !ok {
		return nil, fmt.Errorf("%s/%s: object does not contain %s", typeName, field, field)
	}

	return value, nil
}

// parseRetriever checks if the method has the correct signature and returns the corresponding reflect.Type
// it should take 0 or 1 arguments
// if it has a single argument it should be of type RetrieverInput
// it should return 1 or 2 values
// the first value should be a pointer to the object or a slice of pointers to the object
// the second value should be an error
func parseRetriever(methodTyp reflect.Type) (reflect.Type, error) {
	var typ reflect.Type

	if methodTyp.NumIn() > 1 {
		return typ, fmt.Errorf("method should take 0 or 1 arguments, got %d", methodTyp.NumIn())
	}

	if methodTyp.NumIn() == 1 && methodTyp.In(0) != reflect.TypeOf(RetrieverInput{}) {
		return typ, fmt.Errorf("method should take a single argument of type RetrieverInput, got %s", methodTyp.In(0).String())
	}

	if methodTyp.NumOut() < 1 || methodTyp.NumOut() > 2 {
		return typ, fmt.Errorf("method should return 1 or 2 values, got %d", methodTyp.NumOut())
	}

	firstOut := methodTyp.Out(0)
	if firstOut.Kind() == reflect.Slice {
		if firstOut.Elem().Kind() != reflect.Ptr {
			return typ, fmt.Errorf("first return value should be a slice of pointers, got %s", firstOut.Elem().Kind().String())
		}

		typ = firstOut.Elem().Elem()
	} else if firstOut.Kind() == reflect.Ptr {
		if firstOut.Elem().Kind() == reflect.Slice {
			return typ, fmt.Errorf("first return value should be a pointer to an object, got a pointer to a slice")
		}

		typ = firstOut.Elem()
	} else {
		return typ, fmt.Errorf("first return value should be a pointer or a slice, got %s", firstOut.Kind().String())
	}

	if methodTyp.NumOut() == 2 && methodTyp.Out(1) != reflect.TypeOf((*error)(nil)).Elem() {
		return typ, fmt.Errorf("second return value should be an error, got %s", methodTyp.Out(1).String())
	}

	return typ, nil
}

type retreiversConfig map[string]*retrieversObjectConfig
type retrieversObjectConfig struct {
	primaryKey string
	typ        reflect.Type
	retreivers retrievers
}
type retrievers map[string]reflect.Value

var builtInScalars = map[string]*graphql.Scalar{
	"String":  graphql.String,
	"Int":     graphql.Int,
	"Float":   graphql.Float,
	"Boolean": graphql.Boolean,
	"ID":      graphql.ID,
}
