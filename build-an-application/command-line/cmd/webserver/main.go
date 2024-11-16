package main

import (
	"log"
	"net/http"

	poker "learn-go-with-test/build-an-application"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	//game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	//cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	//cli.PlayPoker()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
