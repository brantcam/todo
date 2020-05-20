package list

// Repo contains all the functionality of todo
type Repo interface {
	GetList() ([]Item, error)
}

// Item is an item in a todo list
type Item struct {
	Todo string
	Done bool
}
