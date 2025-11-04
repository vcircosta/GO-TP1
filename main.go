package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	ID    int
	Name  string
	Email string
}

var lastID int

var contacts = make(map[int]Contact)

var reader = bufio.NewReader(os.Stdin)

func main() {
	setTestContacts()

	for {
		fmt.Println("\n--- Mini CRM 2.0 du Turfu ---")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister tous les contacts")
		fmt.Println("3. Mettre à jour un contact")
		fmt.Println("4. Supprimer un contact")
		fmt.Println("5. Quitter")
		fmt.Print("Choix: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur de lecture :", err)
			continue
		}

		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			handlenewContact()
		case "2":
			listContacts()
		case "3":
			handleUpdateContact()
		case "4":
			handleDeleteContact()
		case "5":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

func setTestContacts() {
	newContact("Alice", "alice@example.com")
	newContact("Bob", "bob@example.com")
}

// ------------------------ 1.Create Contact ------------------------
func handlenewContact() {
	fmt.Print("Nom: ")
	name, _ := reader.ReadString('\n')
	nameValue := strings.TrimSpace(name)

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	emailValue := strings.TrimSpace(email)

	if nameValue == "" || emailValue == "" {
		fmt.Println("Ni le nom ni l'email ne peuvent être vides. Le contact n'a pas pu être ajouté. Veuillez réessayer.")
		return
	}

	newContact(nameValue, emailValue)
	fmt.Println("Contact ajouté avec succès !")

}

func newContact(name, email string) {
	lastID++
	contacts[lastID] = Contact{ID: lastID, Name: name, Email: email}
}

// ------------------------ 2.Show Contact ------------------------
func listContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact")
		return
	}
	fmt.Println("\nListe des contacts :")
	for _, c := range contacts {
		fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Name, c.Email)
	}
}

// ------------------------ 3.Update Contact ------------------------
func handleUpdateContact() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à mettre à jour")
		return
	}

	listContacts()
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

	updateContact(id, name, email)
	fmt.Println("Contact mis à jour avec succès !")
}

func updateContact(id int, name string, email string) {
	contacts[id] = Contact{ID: id, Name: name, Email: email}
}

// ------------------------ 4.Delete Contact ------------------------
func handleDeleteContact() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à supprimer")
		return
	}

	listContacts()
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

	deleteContact(id)
	fmt.Println("Contact supprimé")
}

func deleteContact(id int) {
	delete(contacts, id)
}
