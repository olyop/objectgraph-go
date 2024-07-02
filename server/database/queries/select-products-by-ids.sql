SELECT
	products.product_id,
	products.product_name,
	products.brand_id,
	prices.price_value,
	promotions.promotion_discount,
	promotions.promotion_discount_multiple,
	products_abvs.abv,
	products_volumes.volume,
	products.updated_at,
	products.created_at
FROM
	products
	LEFT JOIN products_prices ON products.product_id = products_prices.product_id
	LEFT JOIN products_promotions ON products.product_id = products_promotions.product_id
	LEFT JOIN products_abvs ON products.product_id = products_abvs.product_id
	LEFT JOIN products_volumes ON products.product_id = products_volumes.product_id
	LEFT JOIN prices ON products_prices.price_id = prices.price_id
	LEFT JOIN promotions ON products_promotions.promotion_id = promotions.promotion_id
WHERE
	products.product_id = ANY ($1);
