package graphqlops

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/graph-gophers/graphql-go"
)

type Handler struct {
	Engine *Engine
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	resolverRequestLocker := new(sync.Map)

	ctx := request.Context()
	ctx = context.WithValue(ctx, EngineContextKey{}, h.Engine)
	ctx = context.WithValue(ctx, ResolverLockerContextKey{}, h.Engine.resolverLocker)
	ctx = context.WithValue(ctx, ResolverRequestLockerContextKey{}, resolverRequestLocker)
	request = request.WithContext(ctx)

	var body graphQLBody
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response := h.Engine.schema.Exec(
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
