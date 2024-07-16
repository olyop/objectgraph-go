package populate

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/olyop/objectgraph-go/data/database"
	"github.com/olyop/objectgraph-go/data/files"
)

func populateUsers(data *files.Data, contactTypes map[string]string) map[string]string {
	var contactsSql strings.Builder
	var personsSql strings.Builder

	contactsParams := make([]any, 0)
	personsParams := make([]any, 0)
	contactsParams = append(contactsParams, contactTypes["email_address"], contactTypes["mobile_number"])
	contactsSql.WriteString("INSERT INTO contacts (contact_type_id, contact_value) VALUES ")
	personsSql.WriteString("INSERT INTO persons (person_first_name, person_last_name, person_dob) VALUES ")

	for i, user := range data.Users {
		firstName := user[0]
		lastName := user[1]
		emailAddress := user[2]
		mobileNumber := generateRandomMobileNumber()
		dob := generateRandomDOB()

		contactsValues := fmt.Sprintf("($1, $%d), ($2, $%d)", len(contactsParams)+1, len(contactsParams)+2)
		personsValues := fmt.Sprintf("($%d, $%d, $%d)", len(personsParams)+1, len(personsParams)+2, len(personsParams)+3)

		var contactsRow string
		var personsRow string
		if i < len(data.Users)-1 {
			contactsRow = fmt.Sprintf("%s, ", contactsValues)
			personsRow = fmt.Sprintf("%s, ", personsValues)
		} else {
			contactsRow = contactsValues
			personsRow = personsValues
		}
		contactsSql.WriteString(contactsRow)
		personsSql.WriteString(personsRow)

		contactsParams = append(contactsParams, emailAddress, mobileNumber)
		personsParams = append(personsParams, firstName, lastName, dob)
	}

	contactsSql.WriteString(" RETURNING contact_id, contact_value")
	personsSql.WriteString(" RETURNING person_id")

	contactsRows, err := database.DB.Query(contactsSql.String(), contactsParams...)
	if err != nil {
		log.Default().Fatal(err)
	}
	log.Default().Println("Populated contacts")

	personsRows, err := database.DB.Query(personsSql.String(), personsParams...)
	if err != nil {
		log.Default().Fatal(err)
	}
	log.Default().Println("Populated persons")

	defer contactsRows.Close()
	defer personsRows.Close()
	contacts := contactsRowsMapper(contactsRows)
	persons := personsRowsMapper(personsRows)
	personsContacts := combinePersonsWithContacts(persons, contacts)

	var usersSql strings.Builder
	var personsContactsSql strings.Builder
	usersParams := make([]any, 0)
	personsContactsParams := make([]any, 0)
	usersSql.WriteString("INSERT INTO users (user_name) VALUES ")
	personsContactsSql.WriteString("INSERT INTO persons_contacts (person_id, contact_id) VALUES ")

	i := 0
	for personID, contactIds := range personsContacts {
		usersValues := fmt.Sprintf("($%d)", i+1)
		personsContactsValues := fmt.Sprintf("($%d, $%d), ($%d, $%d)", len(personsContactsParams)+1, len(personsContactsParams)+2, len(personsContactsParams)+1, len(personsContactsParams)+3)

		var usersRow string
		var personsContactsRow string
		if i < len(personsContacts)-1 {
			usersRow = fmt.Sprintf("%s, ", usersValues)
			personsContactsRow = fmt.Sprintf("%s, ", personsContactsValues)
		} else {
			usersRow = usersValues
			personsContactsRow = personsContactsValues
		}
		usersSql.WriteString(usersRow)
		personsContactsSql.WriteString(personsContactsRow)

		emailAddress := personsContacts[personID][0][1]
		usersParams = append(usersParams, generateUserNameFromEmailAddress(emailAddress))
		personsContactsParams = append(personsContactsParams, personID, contactIds[0][0], contactIds[1][0])

		i++
	}

	usersSql.WriteString(" RETURNING user_id")

	_, err = database.DB.Exec(personsContactsSql.String(), personsContactsParams...)
	if err != nil {
		log.Default().Fatal(err)
	}
	log.Default().Println("Populated persons_contacts")

	usersRows, err := database.DB.Query(usersSql.String(), usersParams...)
	if err != nil {
		log.Default().Fatal(err)
	}
	log.Default().Println("Populated users")

	defer usersRows.Close()
	users := usersRowsMapper(usersRows)

	var usersPersonsSql strings.Builder
	usersPersonsParams := make([]any, 0)
	usersPersonsSql.WriteString("INSERT INTO users_persons (user_id, person_id) VALUES ")

	i = 0
	for userID := range users {
		values := fmt.Sprintf("($%d, $%d)", len(usersPersonsParams)+1, len(usersPersonsParams)+2)

		var row string
		if i < len(users)-1 {
			row = fmt.Sprintf("%s, ", values)
		} else {
			row = values
		}
		usersPersonsSql.WriteString(row)

		usersPersonsParams = append(usersPersonsParams, userID, persons[i])

		i++
	}

	_, err = database.DB.Exec(usersPersonsSql.String(), usersPersonsParams...)
	if err != nil {
		log.Default().Fatal(err)
	}
	log.Default().Println("Populated users persons")

	return users
}

func contactsRowsMapper(rows *sql.Rows) [][2]string {
	contacts := make([][2]string, 0)

	for rows.Next() {
		var contactID string
		var contactValue string

		err := rows.Scan(&contactID, &contactValue)
		if err != nil {
			log.Default().Fatal(err)
		}

		contacts = append(contacts, [2]string{contactID, contactValue})
	}

	return contacts
}

func personsRowsMapper(rows *sql.Rows) []string {
	persons := make([]string, 0)

	for rows.Next() {
		var personID string

		err := rows.Scan(&personID)
		if err != nil {
			log.Default().Fatal(err)
		}

		persons = append(persons, personID)
	}

	return persons
}

func combinePersonsWithContacts(persons []string, contacts [][2]string) map[string][2][2]string {
	personsContacts := make(map[string][2][2]string)

	x := 0
	for _, personID := range persons {
		personsContacts[personID] = [2][2]string{{contacts[x][0], contacts[x][1]}, {contacts[x+1][0], contacts[x+1][1]}}

		x += 2
	}

	return personsContacts
}

func usersRowsMapper(rows *sql.Rows) map[string]string {
	users := make(map[string]string)

	for rows.Next() {
		var userID string

		err := rows.Scan(&userID)
		if err != nil {
			log.Default().Fatal(err)
		}

		users[userID] = userID
	}

	return users
}
