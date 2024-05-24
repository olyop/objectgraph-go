SELECT
	products.product_id,
	products.product_name,
	products.brand_id,
	prices.price_value,
	products_abvs.abv,
	products_volumes.volume,
	products.updated_at,
	products.created_at
FROM
	products
	LEFT JOIN products_prices ON products.product_id = products_prices.product_id
	LEFT JOIN prices ON products_prices.price_id = prices.price_id
	LEFT JOIN products_abvs ON products.product_id = products_abvs.product_id
	LEFT JOIN products_volumes ON products.product_id = products_volumes.product_id
ORDER BY
	products.product_name
LIMIT
	100;
