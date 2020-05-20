package main

import (
	"context"
	"log"
	"todo/store/pg"
)

func main() {
	ctx := context.Background()

	db, err := pg.New(ctx)
	if err != nil {
		log.Fatalf("failed to establish connection to postgres db: %v", err)
	}
}
