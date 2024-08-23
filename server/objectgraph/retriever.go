package objectgraph

import (
	"fmt"
	"reflect"
)

func createRetrievers(config *Configuration) (retreivers, error) {
	retrievers := make(retreivers)

	for typeName, objectConfig := range config.Objects {
		var typ reflect.Type
		retrieversMap := make(map[string]reflect.Value)

		// using reflect, loop over all methods on the retrievers object
		retrieversValue := reflect.ValueOf(objectConfig.Retrievers)
		retrieversTyp := retrieversValue.Type()
		for i := 0; i < retrieversValue.NumMethod(); i++ {
			method := retrieversTyp.Method(i)
			methodVal := retrieversValue.Method(i)
			methodTyp := methodVal.Type()

			retrieverTyp, err := parseRetriever(methodTyp)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", typeName, err)
			}

			if typ == nil {
				typ = retrieverTyp
			} else if typ != retrieverTyp {
				return nil, fmt.Errorf("%s: all methods should return the same type, got %s and %s", typeName, typ.String(), retrieverTyp.String())
			}

			retrieversMap[method.Name] = methodVal
		}

		retrievers[typeName] = &retrieversObject{
			primaryKey: objectConfig.PrimaryKey,
			typ:        typ,
			retreivers: retrieversMap,
		}
	}

	return retrievers, nil
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

type retreivers map[string]*retrieversObject
type retrieversObject struct {
	primaryKey string
	typ        reflect.Type
	retreivers map[string]reflect.Value
}
