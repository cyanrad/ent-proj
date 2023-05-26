package main

import (
	"context"
	"log"
	"main/ent"
	"main/resolver"
	"net/http"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
)

func main() {
	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.Postgres, "postgresql://root:secret@localhost:5433/bank?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// migrate
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(resolver.NewSchema(client))
	http.Handle("/",
		playground.Handler("Coffee", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
