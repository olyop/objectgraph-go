package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/machinebox/graphql"
)

func main() {
	graphQLClient := graphql.NewClient("http://localhost:8080/graphql")

	for i := 0; i < 10; i++ {
		TestGetProductsQuery(graphQLClient)

		time.Sleep(1 * time.Second)
	}
}

func TestGetProductsQuery(client *graphql.Client) {
	file, err := os.ReadFile("test.graphql")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	contents := string(file)
	request := graphql.NewRequest(contents)

	ctx := context.Background()

	err = client.Run(ctx, request, &struct{}{})
	if err != nil {
		log.Fatalf("Error getting products: %v", err)
	}
}
