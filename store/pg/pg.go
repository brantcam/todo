package pg

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// New returns a new instance of Conn
func New(ctx context.Context) (*sql.DB, error) {
	// todo - put these vars in a config that gets passed to this function
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s, dbname=%s sslmode=disable", "localhost", "5432", "todouser", "todopass", "todoname"))
	if err != nil {
		return nil, err
	}

	return db, db.PingContext(ctx)
}
