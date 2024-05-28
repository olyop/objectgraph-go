SELECT
	categories.category_id,
	categories.category_name,
	categories.classification_id,
	categories.updated_at,
	categories.created_at
FROM
	products_categories
	JOIN categories ON products_categories.category_id = categories.category_id
WHERE
	products_categories.product_id = $1;
