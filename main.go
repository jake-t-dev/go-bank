package main

import (
	"log"

	_ "github.com/lib/pq"
)

func main() {
	store, err := NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3333", store)
	server.Run()
}
