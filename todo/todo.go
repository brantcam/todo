package todo

// Ops contains the postgres connection pool
type Ops struct{}

// GetList returns all todo items in the postgres db
func (t *Ops) GetList() ([]ListItem, error) {
	listItems := make([]ListItem, 0)
	return listItems, nil
}
