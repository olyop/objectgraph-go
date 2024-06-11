package populate

import (
	"log"

	"github.com/olyop/graphql-go/data/database"
	"github.com/olyop/graphql-go/data/files"
)

func populateUsers(data *files.Data) map[string]string {
	return map[string]string{}
}

func clearUsers() {
	_, err := database.DB.Exec("DELETE FROM users_persons")
	if err != nil {
		log.Default().Fatal(err)
	}

	_, err = database.DB.Exec("DELETE FROM users")
	if err != nil {
		log.Default().Fatal(err)
	}
	_, err = database.DB.Exec("DELETE FROM persons_contacts")
	if err != nil {
		log.Default().Fatal(err)
	}

	_, err = database.DB.Exec("DELETE FROM persons")
	if err != nil {
		log.Default().Fatal(err)
	}

	_, err = database.DB.Exec("DELETE FROM contacts")
	if err != nil {
		log.Default().Fatal(err)
	}

	_, err = database.DB.Exec("DELETE FROM contact_types")
	if err != nil {
		log.Default().Fatal(err)
	}
}
