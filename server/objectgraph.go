package main

import (
	"embed"
	"os"

	"github.com/olyop/objectgraph/graphql/retrievers"
	"github.com/olyop/objectgraph/objectgraph"
	"github.com/olyop/objectgraph/objectgraph/objectcache"
	"github.com/redis/go-redis/v9"
)

//go:embed graphql/schema/*.graphqls
var schemaFs embed.FS

func createEngineConfiguration() *objectgraph.EngineConfiguration {
	return &objectgraph.EngineConfiguration{
		SourceFiles: schemaFs,
		Cache: &objectcache.Configuration{
			Prefix: os.Getenv("REDIS_PREFIX"),
			Redis: &redis.Options{
				Addr:     os.Getenv("REDIS_URL"),
				Username: os.Getenv("REDIS_USERNAME"),
				Password: os.Getenv("REDIS_PASSWORD"),
			},
		},
		Retrievers: objectgraph.RetrieverMap{
			"Brand": &objectgraph.RetrieverType{
				PrimaryKey: "BrandID",
				Value:      &retrievers.RetrieveBrand{},
			},
			"Category": &objectgraph.RetrieverType{
				PrimaryKey: "CategoryID",
				Value:      &retrievers.RetrieveCategory{},
			},
			"Classification": &objectgraph.RetrieverType{
				PrimaryKey: "ClassificationID",
				Value:      &retrievers.RetrieveClassification{},
			},
			"Contact": &objectgraph.RetrieverType{
				PrimaryKey: "ContactID",
				Value:      &retrievers.RetrieveContact{},
			},
			"Product": &objectgraph.RetrieverType{
				PrimaryKey: "ProductID",
				Value:      &retrievers.RetrieveProduct{},
			},
			"User": &objectgraph.RetrieverType{
				PrimaryKey: "UserID",
				Value:      &retrievers.RetrieveUser{},
			},
		},
	}
}
