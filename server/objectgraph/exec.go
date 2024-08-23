package objectgraph

import (
	"context"
	"time"

	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/structs"
	"github.com/olyop/objectgraph/objectgraph/internal/jsonbuilder"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

func (e *Engine) Exec(ctx context.Context, request *GraphQLRequest) *GraphQLResponse {
	queryAst, err := gqlparser.LoadQuery(e.schemaAst, request.Query)
	if err != nil {
		return handleParseErrors(err)
	}

	// d, _ := json.MarshalIndent(queryAst, "", "  ")
	// os.WriteFile("/home/op/code/objectgraph-go/queryAst.json", d, 0644)

	processor := &queryProcessor{
		engine:   e,
		ctx:      ctx,
		request:  request,
		queryAst: queryAst,
		data:     jsonbuilder.New(),
	}

	return processor.exec()
}

type queryProcessor struct {
	engine   *Engine
	ctx      context.Context
	request  *GraphQLRequest
	queryAst *ast.QueryDocument
	data     *jsonbuilder.JSONBuilder
	errors   []*GraphQLError
}

func (qp *queryProcessor) exec() *GraphQLResponse {
	opAst, err := qp.getOperation(qp.queryAst.Operations)
	if err != nil {
		return qp.exitError(err)
	}

	if opAst.Operation == ast.Query {
		err = qp.execQueryOperation(opAst)
		if err != nil {
			return qp.exitError(err)
		}

		return qp.exit()
	} else if opAst.Operation == ast.Mutation {
		return qp.exitError(fmt.Errorf("mutation is not supported"))
	} else if opAst.Operation == ast.Subscription {
		return qp.exitError(fmt.Errorf("subscription is not supported"))
	} else {
		return qp.exitError(fmt.Errorf("operation type is not supported"))
	}
}

func (qp *queryProcessor) execQueryOperation(opAst *ast.OperationDefinition) error {
	qp.execSelectionSet(opAst.SelectionSet)

	return nil
}

func (qp *queryProcessor) execSelectionSet(selSetAst ast.SelectionSet) {
	jsonObj := make(jsonbuilder.JSONObject, len(selSetAst))

	// dataSelections := make([]DataSelection, len(selSetAst))

	for i, selAst := range selSetAst {
		fieldAst := selAst.(*ast.Field)

		cachedJsonVal := qp.getCachedValue(fieldAst)

		jsonObj[i] = jsonbuilder.JSONPair{
			Key:   fieldAst.Name,
			Value: cachedJsonVal,
		}
	}

	qp.data.WriteObject(jsonObj)
}

func (qp *queryProcessor) getCachedValue(fieldAst *ast.Field) *jsonbuilder.JSONValue {
	retrieveConfig := parseRetrieveConfig(fieldAst)

	typeName := retrieveConfig.objDef.TypeName
	productID := fieldAst.Arguments.ForName("productID").Value.Raw
	valueKey := "object"

	val, _, _ := qp.engine.objectcache.Get(typeName, productID, valueKey, time.Minute, qp.engine.retreivers[typeName].typ)

	return &jsonbuilder.JSONValue{
		JSONType: 1,
		Object:   convertObjectToJSONPair(val, fieldAst),
	}
}

func (qp *queryProcessor) execRetriever(objectType string, funcName string, input RetrieverInput) (any, error) {
	retriever := qp.engine.retreivers[objectType].retreivers[funcName]
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
	}

	return structs.Map(resultValue[0].Interface()), nil
}

func (qp *queryProcessor) getOperation(opList ast.OperationList) (*ast.OperationDefinition, error) {
	if len(opList) != 1 {
		return nil, fmt.Errorf("only one operation is allowed")
	}

	return opList[0], nil
}

func (qp *queryProcessor) getDataSelection(fieldAst *ast.Field) DataSelection {
	return DataSelection{}
}

func (qp *queryProcessor) exit() *GraphQLResponse {
	// if any error clear the data
	if len(qp.errors) > 0 {
		qp.data.Reset()
	}

	return &GraphQLResponse{
		Data:   qp.data.Bytes(),
		Errors: qp.errors,
	}
}

func (qp *queryProcessor) exitError(err error) *GraphQLResponse {
	qp.errors = append(qp.errors, &GraphQLError{
		Message: err.Error(),
	})

	return qp.exit()
}

func parseRetrieveConfig(fieldDefAst *ast.Field) retrieveConfig {
	dirAst := fieldDefAst.Definition.Directives.ForName("retrieve")

	keyArgAst := dirAst.Arguments.ForName("key")
	objDef := parseObjectDefinition(keyArgAst.Value.Raw)

	argsArgAst := dirAst.Arguments.ForName("args")
	args := make(map[string]string)
	for _, arg := range argsArgAst.Value.Children {
		valueSplit := strings.Split(arg.Value.Raw, "=")
		args[valueSplit[0]] = valueSplit[1]
	}

	return retrieveConfig{
		objDef: &objDef,
		args:   args,
	}
}

type retrieveConfig struct {
	objDef *objectDefinition
	args   map[string]string
}

func parseObjectDefinition(value string) objectDefinition {
	s := strings.Split(value, "/")
	typeName := s[0]
	retrieverKey := s[1]

	return objectDefinition{
		TypeName:     typeName,
		RetrieverKey: retrieverKey,
	}
}

type objectDefinition struct {
	TypeName     string
	RetrieverKey string
}

func convertObjectToJSONPair(obj map[string]any, fieldAst *ast.Field) []jsonbuilder.JSONPair {
	pairs := make([]jsonbuilder.JSONPair, len(obj))

	i := 0
	for _, value := range fieldAst.SelectionSet {
		fieldAst := value.(*ast.Field)

		pairs[i] = jsonbuilder.JSONPair{
			Key: fieldAst.Name,
			Value: &jsonbuilder.JSONValue{
				JSONType: 0,
				Data:     obj[fieldAst.Name],
			},
		}

		i++
	}

	return pairs
}

type DataSelection struct {
	TypeName string
}
