package populate

import (
	"fmt"
	"log"
	"math/rand/v2"
	"strings"

	"github.com/olyop/graphql-go/data/database"
	"github.com/olyop/graphql-go/data/import"
)

func populateProducts(data *importdata.Data, classificationsToBrands map[string][]string, brands map[string]string, categories map[string]string) []Product {
	products := generateProducts(data, classificationsToBrands)

	for _, batch := range batchProducts(products) {
		populateProductsBatch(batch, brands, categories)
	}

	return products
}

func generateProducts(data *importdata.Data, classificationsToBrands map[string][]string) []Product {
	products := make([]Product, 0)

	for classification, classificationBrands := range classificationsToBrands {
		types := determineProductTypes(classification)
		categories := data.ClassificationToCategories[classification]

		for _, brand := range classificationBrands {
			for _, category := range categories {
				for _, productType := range types {
					products = append(products, Product{
						name:     fmt.Sprintf("%s %s %s %s", brand, productType.name, category, classification),
						brand:    brand,
						category: category,
						volume:   productType.volume,
						abv:      productType.abv,
						price:    productType.price,
					})
				}
			}
		}
	}

	return products
}

func determineProductTypes(classification string) []Product {
	switch classification {
	case "Beer", "Cider":
		return beerAndCiderProductTypes
	case "Wine":
		return wineProductTypes
	case "Spirits", "Premix":
		return spiritAndPreMixProductTypes
	default:
		panic("Invalid classification")
	}
}

func populateProductsBatch(products []Product, brands map[string]string, categories map[string]string) {
	var sql strings.Builder
	productsParams := make([]interface{}, 0)
	brandsParamsMap := make(map[string]int)

	sql.WriteString("INSERT INTO products (product_name, brand_id) VALUES ")

	for i, product := range products {
		productName := product.name
		productBrandID := brands[product.brand]

		brandParam, ok := brandsParamsMap[productBrandID]
		if !ok {
			brandParam = len(productsParams) + 1
			brandsParamsMap[productBrandID] = brandParam
			productsParams = append(productsParams, productBrandID)
		}

		values := fmt.Sprintf("($%d, $%d)", len(productsParams)+1, brandParam)

		var row string
		if i < len(products)-1 {
			row = fmt.Sprintf("%s,", values)
		} else {
			row = values
		}
		sql.WriteString(row)

		productsParams = append(productsParams, productName)
	}

	sql.WriteString(" RETURNING product_id")

	productsRows, err := database.DB.Query(sql.String(), productsParams...)
	if err != nil {
		panic(err)
	}

	defer productsRows.Close()

	i := 0
	for productsRows.Next() {
		var productID string

		err := productsRows.Scan(&productID)
		if err != nil {
			panic(err)
		}

		products[i].productID = productID

		i++
	}

	var pricesSql strings.Builder
	var promotionsSql strings.Builder
	var productsVolumesSql strings.Builder
	var productsAbvsSql strings.Builder
	var productsCategoriesSql strings.Builder
	var productsPricesSql strings.Builder
	var productsPromotionsSql strings.Builder

	pricesSql.WriteString("INSERT INTO prices (price_value) VALUES ")
	promotionsSql.WriteString("INSERT INTO promotions (promotion_discount, promotion_discount_multiple) VALUES ")
	productsVolumesSql.WriteString("INSERT INTO products_volumes (product_id, volume) VALUES ")
	productsAbvsSql.WriteString("INSERT INTO products_abvs (product_id, abv) VALUES ")
	productsCategoriesSql.WriteString("INSERT INTO products_categories (product_id, category_id) VALUES ")
	productsPricesSql.WriteString("INSERT INTO products_prices (product_id, price_id) VALUES ")
	productsPromotionsSql.WriteString("INSERT INTO products_promotions (product_id, promotion_id) VALUES ")

	productsParamsMap := make(map[string]int)

	productsCategoriesParams := make([]interface{}, 0)
	productsCategoriesProductIDParamsMap := make(map[string]int)
	productsCategoriesCategoryIDParamsMap := make(map[string]int)

	for i, product := range products {
		productIDParam, ok := productsParamsMap[product.productID]
		if !ok {
			productIDParam = len(productsParamsMap) + 1
			productsParamsMap[product.productID] = productIDParam
		}

		categoriesProductIDParam, ok := productsCategoriesProductIDParamsMap[product.productID]
		if !ok {
			categoriesProductIDParam = len(productsCategoriesParams) + 1
			productsCategoriesProductIDParamsMap[product.productID] = categoriesProductIDParam
			productsCategoriesParams = append(productsCategoriesParams, product.productID)
		}

		categoriesCategoryIDParam, ok := productsCategoriesCategoryIDParamsMap[product.category]
		if !ok {
			categoriesCategoryIDParam = len(productsCategoriesParams) + 1
			productsCategoriesCategoryIDParamsMap[product.category] = categoriesCategoryIDParam
			productsCategoriesParams = append(productsCategoriesParams, categories[product.category])
		}

		priceValues := fmt.Sprintf("(%d)", product.price)
		promotionValues := fmt.Sprintf("(%d, %d)", randRange(1, product.price/2), randRange(2, 5))
		volumeValues := fmt.Sprintf("($%d, %d)", productIDParam, product.volume)
		abvValues := fmt.Sprintf("($%d, %d)", productIDParam, product.abv)
		categoryValues := fmt.Sprintf("($%d, $%d)", categoriesProductIDParam, categoriesCategoryIDParam)

		var priceRow string
		var promotionRow string
		var volumeRow string
		var abvRow string
		var categoryRow string
		if i < len(products)-1 {
			priceRow = fmt.Sprintf("%s,", priceValues)
			promotionRow = fmt.Sprintf("%s,", promotionValues)
			volumeRow = fmt.Sprintf("%s,", volumeValues)
			abvRow = fmt.Sprintf("%s,", abvValues)
			categoryRow = fmt.Sprintf("%s,", categoryValues)
		} else {
			priceRow = priceValues
			promotionRow = promotionValues
			volumeRow = volumeValues
			abvRow = abvValues
			categoryRow = categoryValues
		}
		pricesSql.WriteString(priceRow)
		promotionsSql.WriteString(promotionRow)
		productsVolumesSql.WriteString(volumeRow)
		productsAbvsSql.WriteString(abvRow)
		productsCategoriesSql.WriteString(categoryRow)
	}

	pricesSql.WriteString(" RETURNING price_id")
	promotionsSql.WriteString(" RETURNING promotion_id")

	_, err = database.DB.Exec(productsCategoriesSql.String(), productsCategoriesParams...)
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec(productsVolumesSql.String(), convertSetToArr(productsParamsMap)...)
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec(productsAbvsSql.String(), convertSetToArr(productsParamsMap)...)
	if err != nil {
		panic(err)
	}

	pricesRows, err := database.DB.Query(pricesSql.String())
	if err != nil {
		panic(err)
	}

	promotionsRows, err := database.DB.Query(promotionsSql.String())
	if err != nil {
		panic(err)
	}

	defer pricesRows.Close()
	defer promotionsRows.Close()

	j := 0
	for pricesRows.Next() {
		var priceID string

		err := pricesRows.Scan(&priceID)
		if err != nil {
			panic(err)
		}

		productPricesValues := fmt.Sprintf("('%s','%s')", products[j].productID, priceID)

		var productPricesRow string
		if j < len(products)-1 {
			productPricesRow = fmt.Sprintf("%s,", productPricesValues)
		} else {
			productPricesRow = productPricesValues
		}
		productsPricesSql.WriteString(productPricesRow)

		j++
	}

	k := 0
	for promotionsRows.Next() {
		var promotionID string

		err := promotionsRows.Scan(&promotionID)
		if err != nil {
			panic(err)
		}

		productPromotionsValues := fmt.Sprintf("('%s','%s')", products[k].productID, promotionID)

		var productPromotionsRow string
		if k < len(products)-1 {
			productPromotionsRow = fmt.Sprintf("%s,", productPromotionsValues)
		} else {
			productPromotionsRow = productPromotionsValues
		}
		productsPromotionsSql.WriteString(productPromotionsRow)

		k++
	}

	_, err = database.DB.Exec(productsPricesSql.String())
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec(productsPromotionsSql.String())
	if err != nil {
		panic(err)
	}

	log.Printf("Populated %d products", len(products))
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func clearProducts() {
	_, err := database.DB.Exec("DELETE FROM products_prices")
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec("DELETE FROM products_promotions")
	if err != nil {
		panic(err)
	}

	_, err = database.DB.Exec("DELETE FROM prices")
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
