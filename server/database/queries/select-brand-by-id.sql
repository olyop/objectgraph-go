SELECT
	brands.brand_id,
	brands.brand_name,
	brands.updated_at,
	brands.created_at
FROM
	brands
WHERE
	brands.brand_id = $1;
