package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type JSONStore struct {
	filePath string
	contacts map[int]*Contact
	nextID   int
	mu       sync.RWMutex
}

func NewJSONStore(filePath string) (*JSONStore, error) {
	store := &JSONStore{
		filePath: filePath,
		contacts: make(map[int]*Contact),
		nextID:   1,
	}

	if err := store.load(); err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("erreur lors du chargement du fichier: %w", err)
		}
	}

	return store, nil
}

type jsonData struct {
	Contacts map[int]*Contact `json:"contacts"`
	NextID   int              `json:"nextID"`
}

func (js *JSONStore) load() error {
	data, err := os.ReadFile(js.filePath)
	if err != nil {
		return err
	}

	var jd jsonData
	if err := json.Unmarshal(data, &jd); err != nil {
		return fmt.Errorf("erreur de désérialisation: %w", err)
	}

	js.contacts = jd.Contacts
	if js.contacts == nil {
		js.contacts = make(map[int]*Contact)
	}
	js.nextID = jd.NextID
	if js.nextID < 1 {
		js.nextID = 1
	}

	return nil
}

func (js *JSONStore) save() error {
	jd := jsonData{
		Contacts: js.contacts,
		NextID:   js.nextID,
	}

	data, err := json.MarshalIndent(jd, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur d'encodage JSON: %w", err)
	}

	if err := os.WriteFile(js.filePath, data, 0644); err != nil {
		return fmt.Errorf("erreur d'écriture du fichier: %w", err)
	}

	return nil
}

func (js *JSONStore) Add(contact *Contact) error {
	js.mu.Lock()
	defer js.mu.Unlock()

	contact.ID = js.nextID
	js.contacts[contact.ID] = contact
	js.nextID++

	return js.save()
}

func (js *JSONStore) GetAll() ([]*Contact, error) {
	js.mu.RLock()
	defer js.mu.RUnlock()

	all := []*Contact{}
	for _, c := range js.contacts {
		all = append(all, c)
	}
	return all, nil
}

func (js *JSONStore) GetById(id int) (*Contact, error) {
	js.mu.RLock()
	defer js.mu.RUnlock()

	c, ok := js.contacts[id]
	if !ok {
		return nil, fmt.Errorf("Contact avec l'ID %d non trouvé", id)
	}
	return c, nil
}

func (js *JSONStore) Update(id int, newName, newEmail string) error {
	js.mu.Lock()
	defer js.mu.Unlock()

	c, ok := js.contacts[id]
	if !ok {
		return fmt.Errorf("Contact avec l'ID %d non trouvé", id)
	}

	c.Name = newName
	c.Email = newEmail

	return js.save()
}

func (js *JSONStore) Delete(id int) error {
	js.mu.Lock()
	defer js.mu.Unlock()

	_, ok := js.contacts[id]
	if !ok {
		return fmt.Errorf("Contact avec l'ID %d non trouvé", id)
	}

	delete(js.contacts, id)

	return js.save()
}
