package main

import (
	"github.com/elastic/go-ucfg"
	"github.com/khorevaa/r2gitsync/internal/bl"
	"github.com/khorevaa/r2gitsync/internal/di"
	"github.com/khorevaa/r2gitsync/internal/io/http"
	"github.com/khorevaa/r2gitsync/internal/services/datastore"
)

func main() {

	cfg := &ucfg.Config{}

	database, err := datastore.New(cfg)

	if err != nil {
		panic(err)
	}

	deps := di.New(database)

	bLogic := bl.NewBL(deps)

	server := http.New(bLogic)

	server.Run("localhost:4000")
}
