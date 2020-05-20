package pg

import (
	"context"
	"database/sql"
	"fmt"
	"todo/config"

	"github.com/gchaincl/dotsql"
	_ "github.com/lib/pq"
)

// Conn contains the sql queries packaged by dotsql and the postgres connection pool
type Conn struct {
	Queries *dotsql.DotSql
	Db      *sql.DB
}

// New returns a new instance of Conn
func New(ctx context.Context, c *config.Config) (*Conn, error) {
	// get queries from sql file
	queries, err := dotsql.LoadFromFile("./store/sql/todo.sql")
	if err != nil {
		return nil, err
	}
	// todo - put these vars in a config that gets passed to this function
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s, dbname=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName))
	if err != nil {
		return nil, err
	}

	return &Conn{Queries: queries, Db: db}, db.PingContext(ctx)
}
