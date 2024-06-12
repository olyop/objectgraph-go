SELECT
	contacts.contact_id,
	contacts.contact_value,
	contact_types.contact_type_name,
	persons_contacts.updated_at,
	persons_contacts.created_at
FROM
	persons_contacts
	JOIN users_persons ON persons_contacts.person_id = users_persons.person_id
	JOIN contacts ON persons_contacts.contact_id = contacts.contact_id
	JOIN contact_types ON contacts.contact_type_id = contact_types.contact_type_id
WHERE
	users_persons.user_id = $1;
