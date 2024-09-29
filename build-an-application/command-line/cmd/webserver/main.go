package main

import (
	"log"
	"net/http"
	"os"

	poker "github.com/d-shimizu/learn-go-with-test/build-an-application"
)

const dbFileName = "game.db.json"

func main() {
	// OpenFile is used to open a file for reading, writing, or both.
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	// FileSystemPlayerStore is a struct that contains a ReadWriteSeeker
	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	// NewPlayerServer returns a new PlayerServer with a FileSystemPlayerStore
	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}