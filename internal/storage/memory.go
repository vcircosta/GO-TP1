package storage

import "fmt"

type MemoryStore struct {
	contacts map[int]*Contact
	nextID   int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
}

func (ms *MemoryStore) Add(contact *Contact) error {
	contact.ID = ms.nextID
	ms.contacts[contact.ID] = contact
	ms.nextID++
	return nil
}

func (ms *MemoryStore) GetAll() ([]*Contact, error) {
	all := []*Contact{}
	for _, c := range ms.contacts {
		all = append(all, c)
	}
	return all, nil
}

func (ms *MemoryStore) GetById(id int) (*Contact, error) {
	c, ok := ms.contacts[id]
	if !ok {
		return nil, fmt.Errorf("Contact avec l'ID %d non trouvé", id)
	}
	return c, nil
}

func (ms *MemoryStore) Update(id int, newName, newEmail string) error {
	c, ok := ms.contacts[id]
	if !ok {
		return fmt.Errorf("Contact avec l'ID %d non trouvé", id)
	}
	c.Name = newName
	c.Email = newEmail
	return nil
}

func (ms *MemoryStore) Delete(id int) error {
	_, ok := ms.contacts[id]
	if !ok {
		return fmt.Errorf("Contact avec l'ID %d non trouvé", id)
	}
	delete(ms.contacts, id)
	return nil
}
