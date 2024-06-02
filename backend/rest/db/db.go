package db

import (
	"database/sql"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func NewTursoStorage(url string) (*sql.DB, error) {
	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Turso DB: Successfully connected!")

	return db, nil
}
