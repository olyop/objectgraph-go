package objectgraph

import (
	"encoding/json"
)

type GraphQLRequest struct {
	Query         string         `json:"query"`
	OperationName string         `json:"operationName"`
	Variables     map[string]any `json:"variables"`
}

type GraphQLResponse struct {
	Data       json.RawMessage `json:"data"`
	Errors     []*GraphQLError `json:"errors,omitempty"`
	Extensions map[string]any  `json:"extensions,omitempty"`
}

func (r *GraphQLResponse) HasErrors() bool {
	return len(r.Errors) > 0
}

type GraphQLError struct {
	Message    string            `json:"message"`
	Locations  []GraphQLLocation `json:"locations,omitempty"`
	Path       []any             `json:"path,omitempty"`
	Extensions map[string]any    `json:"extensions,omitempty"`
}

type GraphQLLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}
