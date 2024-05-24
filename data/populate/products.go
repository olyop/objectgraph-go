package populate

import (
	"fmt"
	"strings"

	"github.com/olyop/graphql-go/data/database"
	"github.com/olyop/graphql-go/data/import"
)

func populateProducts(data *importdata.Data, classificationsToBrands map[string][]string, brands map[string]string, categories map[string]string) {
	splitProducts := batchProducts(generateProducts(data, classificationsToBrands))

	for _, products := range splitProducts {
		populateProductsBatch(products, brands, categories)
	}
}

func generateProducts(data *importdata.Data, classificationsToBrands map[string][]string) []Product {
	products := make([]Product, 0)

	for classification, classificationBrands := range classificationsToBrands {
		for _, brand := range classificationBrands {
			categories := data.ClassificationToCategories[classification]

			for _, category := range categories {
				types := make([]Product, 0)

				switch classification {
				case "Beer", "Cider":
					types = beerAndCiderProductTypes
				case "Wine":
					types = wineProductTypes
				case "Spirits", "Premix":
					types = spiritAndPreMixProductTypes
				}

				for _, productType := range types {
					productName := fmt.Sprintf("%s %s %s %s", brand, productType.name, category, classification)

					products = append(products, Product{
						name:     productName,
						brand:    brand,
						category: category,
						volume:   productType.volume,
						price:    productType.price,
						abv:      productType.abv,
					})
				}
			}
		}
	}

	return products
}

func populateProductsBatch(products []Product, brands map[string]string, categories map[string]string) {
	var sql strings.Builder
	params := make([]interface{}, 0)

	sql.WriteString("INSERT INTO products (product_name, brand_id) VALUES ")

	for i, product := range products {
		paramsLength := len(params)

		row := fmt.Sprintf("($%d, $%d)", paramsLength+1, paramsLength+2)

		if i < len(products)-1 {
			sql.WriteString(fmt.Sprintf("%s,", row))
		} else {
			sql.WriteString(row)
		}

		params = append(params, product.name, brands[product.brand])
	}

	sql.WriteString(" RETURNING product_id")

	rows, err := database.DB.Query(sql.String(), params...)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	i := 0
	for rows.Next() {
		var productID string

		err := rows.Scan(&productID)
		if err != nil {
			panic(err)
		}

		products[i].productID = productID

		fmt.Printf("Added product: %s, %s\n", productID, products[i].name)

		i++
	}

	var pricesSql strings.Builder
	var productsCategoriesSql strings.Builder
	var productsPricesSql strings.Builder
	var productsVolumesSql strings.Builder
	var productsAbvsSql strings.Builder

	pricesSql.WriteString("INSERT INTO prices (price_value) VALUES ")
	productsCategoriesSql.WriteString("INSERT INTO products_categories (product_id, category_id) VALUES ")
	productsPricesSql.WriteString("INSERT INTO products_prices (product_id, price_id) VALUES ")
	productsVolumesSql.WriteString("INSERT INTO products_volumes (product_id, volume) VALUES ")
	productsAbvsSql.WriteString("INSERT INTO products_abvs (product_id, abv) VALUES ")

	for i, product := range products {
		priceRow := fmt.Sprintf("(%d)", product.price)
		categoryRow := fmt.Sprintf("('%s', '%s')", product.productID, categories[product.category])
		volumeRow := fmt.Sprintf("('%s', %d)", product.productID, product.volume)
		abvRow := fmt.Sprintf("('%s', %d)", product.productID, product.abv)

		if i < len(products)-1 {
			pricesSql.WriteString(fmt.Sprintf("%s,", priceRow))
			productsCategoriesSql.WriteString(fmt.Sprintf("%s,", categoryRow))
			productsVolumesSql.WriteString(fmt.Sprintf("%s,", volumeRow))
			productsAbvsSql.WriteString(fmt.Sprintf("%s,", abvRow))
		} else {
			pricesSql.WriteString(priceRow)
			productsCategoriesSql.WriteString(categoryRow)
			productsVolumesSql.WriteString(volumeRow)
			productsAbvsSql.WriteString(abvRow)
		}
	}

	pricesSql.WriteString(" RETURNING price_id")

	pricesRows, err := database.DB.Query(pricesSql.String())
	if err != nil {
		panic(err)
	}

	defer pricesRows.Close()

	j := 0
	for pricesRows.Next() {
		var priceID string

		err := pricesRows.Scan(&priceID)
		if err != nil {
			panic(err)
		}

		productPricesRow := fmt.Sprintf("('%s', '%s')", products[j].productID, priceID)

		if j < len(products)-1 {
			productsPricesSql.WriteString(fmt.Sprintf("%s,", productPricesRow))
		} else {
			productsPricesSql.WriteString(productPricesRow)
		}

		j++
	}

	_, err = database.DB.Exec(productsCategoriesSql.String())
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec(productsPricesSql.String())
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec(productsVolumesSql.String())
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec(productsAbvsSql.String())
	if err != nil {
		panic(err)
	}
}

func clearProducts() {
	_, err := database.DB.Exec("DELETE FROM products_prices")
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec("DELETE FROM prices")
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec("DELETE FROM products_categories")
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec("DELETE FROM products_volumes")
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec("DELETE FROM products_abvs")
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec("DELETE FROM products")
	if err != nil {
		panic(err)
	}
}

type Product struct {
	productID string
	name      string
	brand     string
	category  string
	volume    int
	price     int
	abv       int
}
