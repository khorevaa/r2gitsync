//go:build ignore
// +build ignore

package main

import (
	"log"
	
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	if err := entc.Generate("../../internal/schema", &gen.Config{
		Target:  "../../internal/services/db",
		Package: "github.com/khorevaa/r2gitsync/internal/services/db",
	}); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
