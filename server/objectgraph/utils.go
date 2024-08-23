package objectgraph

import (
	"github.com/vektah/gqlparser/gqlerror"
)

func handleParseErrors(errList gqlerror.List) *GraphQLResponse {
	graphqlErrors := make([]*GraphQLError, len(errList))

	for i, err := range errList {
		graphqlErrors[i] = &GraphQLError{
			Message: err.Message,
		}
	}

	return &GraphQLResponse{
		Errors: graphqlErrors,
	}
}

func handleExecErrors(err error) *GraphQLResponse {
	return &GraphQLResponse{
		Errors: []*GraphQLError{
			{
				Message: err.Error(),
			},
		},
	}
}
