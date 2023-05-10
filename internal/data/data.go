package data

import (
	"database/sql"
	"log"
	"sync"
)

var (
	data *Data
	once sync.Once
)

type Data struct {
	DB *sql.DB
}

func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connected to database successfully!")

	err = MakeMigration(db)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Migrated database successfully!")

	data = &Data{
		DB: db,
	}
}

func New() *Data {
	once.Do(initDB)

	return data
}

func Close() *Data {
	err := data.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
	return data
}
