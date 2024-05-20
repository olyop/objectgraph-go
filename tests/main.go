package main

import (
	"context"
	"log"
	"time"

	"github.com/machinebox/graphql"
)

func main() {
	graphQLClient := graphql.NewClient("http://localhost:8080/graphql")

	for i := 0; i < 100; i++ {
		TestGetProductsQuery(graphQLClient)

		time.Sleep(1 * time.Second)
	}
}

func TestGetProductsQuery(client *graphql.Client) {
	req := graphql.NewRequest(`
		query {
			getProducts {
				productID
				name
				price
				volume
				abv
				createdAt
				brand {
					brandID
					name
					createdAt
				}
				categories {
					categoryID
					name
					createdAt
				}
			}
		}
	`)

	ctx := context.Background()

	var response GetProductsResponse
	err := client.Run(ctx, req, &response)
	if err != nil {
		log.Fatalf("Error getting products: %v", err)
	}

	if len(response.GetProducts) == 0 {
		log.Fatal("No products found")
	}

	for _, product := range response.GetProducts {
		if product.ProductID == "" {
			log.Fatal("ProductID should not be empty")
		}
		if product.Name == "" {
			log.Fatal("Product name should not be empty")
		}
		if product.Price <= 0 {
			log.Fatal("Price should be greater than 0")
		}
		if product.Brand.BrandID == "" {
			log.Fatal("BrandID should not be empty")
		}
		if product.Brand.Name == "" {
			log.Fatal("Brand name should not be empty")
		}
		for _, category := range product.Categories {
			if category.CategoryID == "" {
				log.Fatal("CategoryID should not be empty")
			}
			if category.Name == "" {
				log.Fatal("Category name should not be empty")
			}
		}
	}
}

type Product struct {
	ProductID  string     `json:"productID"`
	Name       string     `json:"name"`
	Price      int        `json:"price"`
	Volume     int        `json:"volume"`
	Abv        float64    `json:"abv"`
	CreatedAt  int64      `json:"createdAt"`
	Brand      Brand      `json:"brand"`
	Categories []Category `json:"categories"`
}

type Brand struct {
	BrandID   string `json:"brandID"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"createdAt"`
}

type Category struct {
	CategoryID string `json:"categoryID"`
	Name       string `json:"name"`
	CreatedAt  int64  `json:"createdAt"`
}

type GetProductsResponse struct {
	GetProducts []Product `json:"getProducts"`
}
