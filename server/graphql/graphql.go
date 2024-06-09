package graphql

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/olyop/graphql-go/server/engine"
	"github.com/olyop/graphql-go/server/graphql/resolvers"
)

func CreateHandler() http.HandlerFunc {
	engine.RegisterRetrievers(retreiversMap)

	engine.RegisterCacheDurations(map[string]time.Duration{
		"catalog": 15 * time.Minute,
	})

	schemaString, err := ReadSchema("./schema")
	if err != nil {
		log.Fatal(err)
	}

	options := []graphql.SchemaOpt{
		graphql.MaxParallelism(20),
	}

	schema := graphql.MustParseSchema(schemaString, &resolvers.Resolver{}, options...)

	resolverMutexMap := new(sync.Map)

	return func(writer http.ResponseWriter, request *http.Request) {
		requestMutexMap := new(sync.Map)

		request = request.WithContext(
			context.WithValue(
				context.WithValue(
					request.Context(),
					engine.ResolverRetrieveMutexContextKey{},
					resolverMutexMap,
				),
				engine.ResolverRequestMutexContextKey{},
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
