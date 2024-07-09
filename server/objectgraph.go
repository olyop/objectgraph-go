package main

import (
	"embed"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/olyop/objectgraph/graphql/retrievers"
	"github.com/olyop/objectgraph/objectgraph"
	"github.com/redis/go-redis/v9"
)

//go:embed graphql/schema/*.graphqls
var schemaFs embed.FS

func engineConfiguration() *objectgraph.Configuration {
	return &objectgraph.Configuration{
		SourceFiles: schemaFs,
		CachePrefix: os.Getenv("REDIS_PREFIX"),
		CacheRedis: &redis.Options{
			Addr:     os.Getenv("REDIS_URL"),
			Username: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		Objects: objectgraph.ConfigurationObjects{
			"Brand": &objectgraph.ConfigurationObject{
				PrimaryKey: "BrandID",
				Retrievers: &retrievers.RetrieveBrand{},
			},
			"Category": &objectgraph.ConfigurationObject{
				PrimaryKey: "CategoryID",
				Retrievers: &retrievers.RetrieveCategory{},
			},
			"Classification": &objectgraph.ConfigurationObject{
				PrimaryKey: "ClassificationID",
				Retrievers: &retrievers.RetrieveClassification{},
			},
			"Contact": &objectgraph.ConfigurationObject{
				PrimaryKey: "ContactID",
				Retrievers: &retrievers.RetrieveContact{},
			},
			"Product": &objectgraph.ConfigurationObject{
				PrimaryKey: "ProductID",
				Retrievers: &retrievers.RetrieveProduct{},
			},
			"User": &objectgraph.ConfigurationObject{
				PrimaryKey: "UserID",
				Retrievers: &retrievers.RetrieveUser{},
			},
		},
		Scalars: map[string]*graphql.Scalar{
			"UUID":      uuidScalar,
			"Timestamp": timestampScalar,
		},
	}
}
