package populate

import (
	"database/sql"
	"log"

	"github.com/olyop/graphql-go/data/database"
)

func clearTables() {
	tx, err := database.DB.Begin()
	if err != nil {
		log.Default().Fatal(err)
	}

	deleteTable(tx, "products_promotions")
	deleteTable(tx, "products_prices")
	deleteTable(tx, "products_categories")
	deleteTable(tx, "products_volumes")
	deleteTable(tx, "products_abvs")
	deleteTable(tx, "products")
	deleteTable(tx, "prices")
	deleteTable(tx, "promotions")
	deleteTable(tx, "brands")
	deleteTable(tx, "categories")
	deleteTable(tx, "classifications")
	deleteTable(tx, "users_persons")
	deleteTable(tx, "users")
	deleteTable(tx, "persons_contacts")
	deleteTable(tx, "contacts")
	deleteTable(tx, "contact_types")

	err = tx.Commit()
	if err != nil {
		log.Default().Fatal(err)
	}
}

func deleteTable(tx *sql.Tx, tableName string) {
	_, err := tx.Exec("DELETE FROM " + tableName)
	if err != nil {
		log.Default().Fatal(err)
	}

	log.Default().Printf("Deleted %s\n", tableName)
}
