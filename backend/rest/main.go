package main

import (
	"fmt"
	"log"
)

func main() {

	database := InitConfig().DBName
	token := InitConfig().DBToken

	databaseURL := fmt.Sprintf("libsql://%s.turso.io?authToken=%s", database, token)

	tursoStorage := NewTursoStorage(databaseURL)

	db, err := tursoStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)
	api := NewAPIServer(":8080", store)
	api.Run()
}
