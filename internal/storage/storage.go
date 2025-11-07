package storage

import "fmt"

type Contact struct {
	ID    int    `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"size:100;not null"`
	Email string `gorm:"size:100;not null"`
}

type Storer interface {
	Add(contact *Contact) error
	GetAll() ([]*Contact, error)
	GetById(id int) (*Contact, error)
	Update(id int, newName string, newEmail string) error
	Delete(id int) error
}

var ErrContactNoFound = func(id int) error { return fmt.Errorf("Contact avec l'ID %d non trouvé.", id) }
