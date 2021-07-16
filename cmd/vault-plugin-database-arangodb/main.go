package main

import (
	"log"
	"os"

	arangodb "github.com/AnatoliyHanukaev/vault-plugin-database-arangodb"
	"github.com/hashicorp/vault/sdk/database/dbplugin/v5"
)

func main() {
	if err := Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func Run() error {
	db, err := arangodb.New()
	if err != nil {
		return err
	}

	dbplugin.Serve(db.(dbplugin.Database))
	return nil
}
