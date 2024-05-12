package main

import "log"

func main() {
	databaseURL := "libsql://[DATABASE].turso.io?authToken=[TOKEN]"

	tursoStorage := NewTursoStorage(databaseURL)

	db, err := tursoStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)
	api := NewAPIServer(":8080", store)
	api.Run()
}