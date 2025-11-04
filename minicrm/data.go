package minicrm

type Contact struct {
	ID    int
	Name  string
	Email string
}

var lastID int

var contacts = make(map[int]Contact)
