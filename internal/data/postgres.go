package data

import (
	"database/sql"
	"gestorpasswordapi/database"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	uri := os.Getenv("DATABASE_URI")
	return sql.Open("postgres", uri)
}

func MakeMigration(db *sql.DB) error {
	var models = &database.Models
	b, err := models.ReadFile("models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
