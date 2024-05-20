SELECT
	products.product_id,
	products.name,
	products.brand_id,
	products_prices.value,
	products_abv.abv,
	products_volumes.volume,
	products.created_at
FROM
	products
	JOIN products_prices ON products.product_id = products_prices.product_id
	LEFT JOIN products_abv ON products.product_id = products_abv.product_id
	LEFT JOIN products_volumes ON products.product_id = products_volumes.product_id
WHERE
	products.product_id = $1
LIMIT
	100;
