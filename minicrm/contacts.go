package minicrm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// ------------------------ Initialisation ------------------------
func SetTestContacts() {
	NewContact("Alice", "alice@example.com")
	NewContact("Bob", "bob@example.com")
}

// ------------------------ 1. Create Contact ------------------------
func HandleNewContact() {
	fmt.Print("Nom: ")
	name, _ := reader.ReadString('\n')
	nameValue := strings.TrimSpace(name)

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	emailValue := strings.TrimSpace(email)

	if nameValue == "" || emailValue == "" {
		fmt.Println("Ni le nom ni l'email ne peuvent être vides. Le contact n'a pas pu être ajouté.")
		return
	}

	NewContact(nameValue, emailValue)
	fmt.Println("Contact ajouté avec succès !")
}

func NewContact(name string, email string) {
	lastID++
	contacts[lastID] = Contact{ID: lastID, Name: name, Email: email}
}

// ------------------------ 2. Show Contact ------------------------
func ListContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact")
		return
	}
	fmt.Println("\nListe des contacts :")
	for _, c := range contacts {
		fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Name, c.Email)
	}
}

// ------------------------ 3. Update Contact ------------------------
func HandleUpdateContact() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à mettre à jour")
		return
	}

	ListContacts()
	fmt.Print("\nID du contact à mettre à jour: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	c, ok := contacts[id]
	if !ok {
		fmt.Println("Aucun contact avec cet ID")
		return
	}

	fmt.Printf("Nom actuel (%s), appuyez sur Entrée pour garder: ", c.Name)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = c.Name
	}

	fmt.Printf("Email actuel (%s), appuyez sur Entrée pour garder: ", c.Email)
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email == "" {
		email = c.Email
	}

	UpdateContact(id, name, email)
	fmt.Println("Contact mis à jour avec succès !")
}

func UpdateContact(id int, name string, email string) {
	contacts[id] = Contact{ID: id, Name: name, Email: email}
}

// ------------------------ 4. Delete Contact ------------------------
func HandleDeleteContact() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à supprimer")
		return
	}

	ListContacts()
	fmt.Print("\nID du contact à supprimer: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	_, ok := contacts[id]
	if !ok {
		fmt.Println("Aucun contact avec cet ID")
		return
	}

	DeleteContact(id)
	fmt.Println("Contact supprimé")
}

func DeleteContact(id int) {
	delete(contacts, id)
}
