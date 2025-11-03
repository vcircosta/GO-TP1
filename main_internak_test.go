package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func resetContacts() {
	contacts = make(map[int]Contact)
}

func TestAddContact(t *testing.T) {
	resetContacts()
	addContact(1, "Alice", "alice@example.com")

	if c, ok := contacts[1]; !ok {
		t.Error("Contact non ajouté")
	} else if c.Name != "Alice" || c.Email != "alice@example.com" {
		t.Error("Les informations du contact sont incorrectes")
	}
}

func TestDeleteContact(t *testing.T) {
	resetContacts()
	addContact(2, "Bob", "bob@example.com")
	delete(contacts, 2)

	if _, ok := contacts[2]; ok {
		t.Error("Le contact n'a pas été supprimé")
	}
}

func TestUpdateContact(t *testing.T) {
	resetContacts()
	addContact(3, "Charlie", "charlie@example.com")

	contacts[3] = Contact{ID: 3, Name: "Charlie Updated", Email: "charlie2@example.com"}

	c, _ := contacts[3]
	if c.Name != "Charlie Updated" || c.Email != "charlie2@example.com" {
		t.Error("Le contact n'a pas été mis à jour correctement")
	}
}

func TestListContacts(t *testing.T) {
	resetContacts()
	addContact(4, "David", "david@example.com")

	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	listContacts()

	w.Close()
	buf.ReadFrom(r)

	os.Stdout = stdout

	output := buf.String()
	if output == "" || !strings.Contains(output, "David") {
		t.Error("Liste des contacts incorrecte")
	}
}
