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

// AddItem will insert new items to the list if they don't already exist
func (t *Ops) AddItem(items []Item) ([]Item, error) {
	for _, item := range items {
		res, err := t.Pg.Queries.ExecContext(context.Background(), t.Pg.Db, "add-item", item.Todo, item.Done)
		if err != nil {
			return nil, err
		}

		rows, err := res.RowsAffected()
		if err != nil {
			return nil, err
		}

		if rows != 1 {
			return nil, fmt.Errorf("did not receive 1 for the number of rows affected, received: %d", rows)
		}
	}

	return t.GetList()
}
