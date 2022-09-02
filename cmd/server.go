package main

import (
	"log"
	"net/http"

	v1 "github.com/alinjie/go-by-example/internal/api/v1"
	"github.com/alinjie/go-by-example/internal/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"))

	if err != nil {
		log.Fatal("failed to connect to db", err)
	}

	storage := storage.New(db)

	if err := storage.Migrate(); err != nil {
		log.Fatal("failed to migrate db", err)
	}

	if err := storage.Seed(); err != nil {
		log.Fatal("failed to seed db", err)
	}

	v1Server := v1.Server(storage)
	http.ListenAndServe("localhost:1234", v1Server)
}
