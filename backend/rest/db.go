package main

import (
	"database/sql"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type TursoStorage struct {
	db *sql.DB
}

func NewTursoStorage(url string) *TursoStorage {
	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Turso")

	return &TursoStorage{
		db: db,
	}
}

func (s *TursoStorage) Init() (*sql.DB, error) {
	//Initialize

	return s.db, nil
}
