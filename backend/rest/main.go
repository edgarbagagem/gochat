package main

import (
	"fmt"
	"log"

	api "github.com/edgarbagagem/gochat/api"
	cfg "github.com/edgarbagagem/gochat/config"
	turso "github.com/edgarbagagem/gochat/db"
)

func main() {
	databaseURL := fmt.Sprintf("libsql://%s.turso.io?authToken=%s", cfg.Envs.DBName, cfg.Envs.DBToken)

	tursoStorage := turso.NewTursoStorage(databaseURL)

	db, err := tursoStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := turso.NewStore(db)
	api := api.NewAPIServer(":8080", store)
	api.Run()
}
