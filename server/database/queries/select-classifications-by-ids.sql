SELECT
	classifications.classification_id,
	classifications.classification_name,
	classifications.updated_at,
	classifications.created_at
FROM
	classifications
WHERE
	classifications.classification_id = ANY ($1);
