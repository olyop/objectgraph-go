SELECT
	products_categories.product_id,
	categories.category_name,
	categories.updated_at,
	categories.created_at
FROM
	products_categories
	JOIN categories ON products_categories.category_id = categories.category_id
WHERE
	products_categories.product_id = $1;
