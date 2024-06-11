package engine

import (
	"context"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"sync"

	"github.com/graph-gophers/graphql-go"
	"github.com/olyop/graphql-go/server/engine/schema"
)

func CreateHandler(schemaFs fs.FS, resolver interface{}) http.HandlerFunc {
	schema := schema.Parse(schemaFs, resolver)

	resolverMutexMap := new(sync.Map)

	return func(writer http.ResponseWriter, request *http.Request) {
		requestMutexMap := new(sync.Map)

		request = request.WithContext(
			context.WithValue(
				context.WithValue(
					request.Context(),
					ResolverRetrieveMutexContextKey{},
					resolverMutexMap,
				),
				ResolverRequestMutexContextKey{},
				requestMutexMap,
			),
		)

		var body graphQLBody
		err := json.NewDecoder(request.Body).Decode(&body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		response := schema.Exec(
			request.Context(),
			body.Query,
			body.OperationName,
			body.Variables,
		)

		for _, err := range response.Errors {
			log.Default().Println("graphql: " + err.Message)
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", determineContentType(response))
		writer.Write(responseJSON)
	}
}

func determineContentType(response *graphql.Response) string {
	if len(response.Errors) > 0 {
		return "application/problem+json"
	}

	return "application/json"
}

type graphQLBody struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}
