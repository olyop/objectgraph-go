package objectgraph

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body GraphQLRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := graphql.Params{
		Context:        r.Context(),
		Schema:         *e.schema,
		OperationName:  body.OperationName,
		RequestString:  body.Query,
		VariableValues: body.Variables,
	}

	result := graphql.Do(params)

	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", determineContentType(result))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func determineContentType(result *graphql.Result) string {
	if result.HasErrors() {
		return "application/problem+json; charset=utf-8"
	}

	return "application/json; charset=utf-8"
}

type GraphQLRequest struct {
	Query         string         `json:"query"`
	OperationName string         `json:"operationName"`
	Variables     map[string]any `json:"variables"`
}
