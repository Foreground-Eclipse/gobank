package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Foreground-Eclipse/gobank/api"
	types "github.com/Foreground-Eclipse/gobank/storage/model"
	"github.com/Foreground-Eclipse/gobank/storage/storage"
)

func seedAccount(store storage.Storage, fname, lname, pw string) *types.Account {
	acc, err := types.NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", acc.Number)

	return acc

}

func seedAccounts(s storage.Storage) {
	seedAccount(s, "white", "bye", "bebe")
}

func main() {
	seed := flag.Bool("seed", true, "seed the db")
	flag.Parse()

	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)

	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	//seed stuff
	if *seed {
		seedAccounts(store)
		fmt.Println("seeding the db")

	}

	server := api.NewAPIServer(":3000", store)
	server.Run()
}
