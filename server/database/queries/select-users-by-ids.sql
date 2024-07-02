SELECT
	users.user_id,
	users.user_name,
	persons.person_first_name,
	persons.person_last_name,
	persons.person_dob,
	users.updated_at,
	users.created_at
FROM
	users
	JOIN users_persons ON users.user_id = users_persons.user_id
	JOIN persons ON users_persons.person_id = persons.person_id
WHERE
	users.user_id = ANY ($1);
