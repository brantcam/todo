package todo

// Repo contains all the functionality of todo
type Repo interface {
	GetList() ([]ListItem, error)
}

// ListItem is an item in a todo list
type ListItem struct {
	ID   int64
	Todo string
}
