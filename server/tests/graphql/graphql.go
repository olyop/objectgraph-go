package graphql

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
)

func RunQuery(query string) {
	log.Default().Print("graphql: running query")

	request := graphql.NewRequest(query)

	response := struct{}{}

	err := client.Run(context.Background(), request, response)
	if err != nil {
		log.Default().Fatalf("graphql: error %v", err)
	}

	log.Default().Print("graphql: query executed")
}
