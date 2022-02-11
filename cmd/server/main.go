package main

import (
	"github.com/elastic/go-ucfg"
	"github.com/khorevaa/r2gitsync/internal/bl"
	"github.com/khorevaa/r2gitsync/internal/di"
	"github.com/khorevaa/r2gitsync/internal/io/http"
	db2 "github.com/khorevaa/r2gitsync/internal/services/datastore"
)

func main() {

	cfg := &ucfg.Config{}

	database, _ := db2.New(cfg)

	deps := di.New(database)

	bLogic := bl.NewBL(deps)

	server := http.New(bLogic)

	server.Run("localhost:4000")
}
