package main

import (
	"context"
	"log"
	"todo/store/pg"
	"todo/todo"
)

func main() {
	ctx := context.Background()

	db, err := pg.New(ctx)
	if err != nil {
		log.Fatalf("failed to establish connection to postgres db: %v", err)
	}

	todo := todo.Ops{Pg: db}
}
