package list

import (
	"context"
	"fmt"
	"todo/store/pg"
)

// Ops contains the postgres connection pool
type Ops struct {
	Pg *pg.Conn
}

// GetList returns all todo items in the postgres db
func (t *Ops) GetList() ([]Item, error) {
	rows, err := t.Pg.Queries.QueryContext(context.Background(), t.Pg.Db, "get-all")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Item, 0)
	for rows.Next() {
		item := Item{}
		if err := rows.Scan(&item.Todo, &item.Done); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occured while reading rows: %s", err)
	}

	return items, nil
}
