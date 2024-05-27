package main

import (
	"fmt"
	"log"

	"github.com/edgarbagagem/gochat/api"
	"github.com/edgarbagagem/gochat/config"
	"github.com/edgarbagagem/gochat/db"
)

func main() {
	databaseURL := fmt.Sprintf("libsql://%s.turso.io?authToken=%s", config.Envs.DBName, config.Envs.DBToken)

	db, err := db.NewTursoStorage(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	api := api.NewAPIServer(":8080", db)
	api.Run()
}
