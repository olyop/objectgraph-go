package graphql

import (
	"github.com/machinebox/graphql"
)

var client *graphql.Client

func InitializeClient() {
	client = graphql.NewClient("https://localhost:8080/graphql")
}
