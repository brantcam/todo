package todo

import (
	"context"
	"todo/store/pg"
)

// Ops contains the postgres connection pool
type Ops struct {
	Pg *pg.Conn
}

// GetList returns all todo items in the postgres db
func (t *Ops) GetList() ([]ListItem, error) {
	rows, err := t.Pg.Queries.QueryContext(context.Background(), t.Pg.Db, "get-all")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]*ListItem, 0)
	for rows.Next() {
		item := &ListItem{}
		if err := rows.Scan(&item.ID, &item.Todo); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	listItems := make([]ListItem, 0)
	return listItems, nil
}
