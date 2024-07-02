SELECT
	categories.category_id,
	categories.category_name,
	categories.classification_id,
	categories.updated_at,
	categories.created_at
FROM
	products_categories
WHERE
	categories.category_id = ANY ($1);
