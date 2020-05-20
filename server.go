package main

import (
	"context"
	"log"
	"net/http"
	"todo/config"
	"todo/list"
	"todo/router"
	"todo/store/pg"
)

func main() {
	c, err := config.LoadConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db, err := pg.New(ctx, c)
	if err != nil {
		log.Fatalf("failed to establish connection to postgres db: %v", err)
	}

	router := router.New(router.Opts{
		List: &list.Ops{Pg: db},
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("listening on port %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
