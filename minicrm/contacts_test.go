package minicrm

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func resetContacts() {
	contacts = make(map[int]Contact)
	lastID = 0
}

func setTestContact() {
	resetContacts()
	NewContact("Alice", "alice@example.com")
}

func TestAddContact(t *testing.T) {
	fmt.Print("AddContact ---------------------------")

	setTestContact()

	if c, ok := contacts[1]; !ok {
		t.Error("Contact non ajouté")
		fmt.Println("-> KO")
	} else if c.Name != "Alice" || c.Email != "alice@example.com" {
		t.Error("Les informations du contact sont incorrectes")
		fmt.Println("-> KO")
	} else {
		fmt.Println("-> OK")
	}
}

func TestDeleteContact(t *testing.T) {
	fmt.Print("DeleteContact ------------------------")

	setTestContact()
	DeleteContact(1)

	if _, ok := contacts[1]; ok {
		t.Error("Le contact n'a pas été supprimé")
		fmt.Println("-> KO")
	} else {
		fmt.Println("-> OK")
	}
}

func TestUpdateContact(t *testing.T) {
	fmt.Print("UpdateContact ------------------------")

	setTestContact()

	UpdateContact(1, "Charlie Updated", "charlie2@example.com")

	c, _ := contacts[1]
	if c.Name != "Charlie Updated" || c.Email != "charlie2@example.com" {
		t.Error("Le contact n'a pas été mis à jour correctement")
		fmt.Println("-> KO")
	} else {
		fmt.Println("-> OK")
	}
}

func TestListContacts(t *testing.T) {
	fmt.Print("ListContacts -------------------------")

	setTestContact()

	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ListContacts()

	w.Close()
	buf.ReadFrom(r)
	os.Stdout = stdout

	output := buf.String()
	if output == "" || !strings.Contains(output, "Alice") {
		t.Error("Liste des contacts incorrecte")
		fmt.Println("-> KO")
	} else {
		fmt.Println("-> OK")
	}
}
