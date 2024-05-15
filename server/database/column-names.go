package database

var columnNames = map[string][]string{
	"brands": {
		"brand_id",
		"name",
	},
	"products": {
		"product_id",
		"name",
		"brand_id",
	},
}
