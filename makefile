generate-db:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

generate-graphql:
	go run -mod=mod ./ent/entc.go
	go run -mod=mod github.com/99designs/gqlgen

generate:
	make generate-db
	make generate-graphql