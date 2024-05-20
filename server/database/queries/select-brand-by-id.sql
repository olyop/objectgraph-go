SELECT
	brands.brand_id,
	brands.name,
	brands.created_at
FROM
	brands
WHERE
	brands.brand_id = $1;
