package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vcircosta/GO-TP1/internal/storage"
)

func Run(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)
	setTestContacts(store)

	for {
		fmt.Println("\n--- Mini CRM 2.1 du Turfu ---")
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

		switch strings.TrimSpace(input) {
		case "1":
			handleNewContact(reader, store)
		case "2":
			listContacts(store)
		case "3":
			handleUpdateContact(reader, store)
		case "4":
			handleDeleteContact(reader, store)
		case "5":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

// ------------------------ Initialisation ------------------------
func setTestContacts(store storage.Storer) {
	store.Add(&storage.Contact{Name: "Alice", Email: "alice@example.com"})
	store.Add(&storage.Contact{Name: "Bob", Email: "bob@example.com"})
}

// ------------------------ Helpers ------------------------
func readLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// ------------------------ 1. Create Contact ------------------------
func handleNewContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Nom: ")
	name := readLine(reader)

	fmt.Print("Email: ")
	email := readLine(reader)

	if name == "" || email == "" {
		fmt.Println("Ni le nom ni l'email ne peuvent être vides. Le contact n'a pas pu être ajouté.")
		return
	}

	err := store.Add(&storage.Contact{Name: name, Email: email})
	if err != nil {
		fmt.Println("Erreur lors de l'ajout du contact :", err)
		return
	}

	fmt.Println("Contact ajouté avec succès !")
}

// ------------------------ 2. Show Contact ------------------------
func listContacts(store storage.Storer) {
	contacts, err := store.GetAll()
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

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
func handleUpdateContact(reader *bufio.Reader, store storage.Storer) {
	contacts, _ := store.GetAll()
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à mettre à jour")
		return
	}

	listContacts(store)

	fmt.Print("\nID du contact à mettre à jour: ")
	id, _ := strconv.Atoi(readLine(reader))

	contact, err := store.GetById(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Nom actuel (%s), appuyez sur Entrée pour garder: ", contact.Name)
	name := readLine(reader)
	if name == "" {
		name = contact.Name
	}

	fmt.Printf("Email actuel (%s), appuyez sur Entrée pour garder: ", contact.Email)
	email := readLine(reader)
	if email == "" {
		email = contact.Email
	}

	err = store.Update(id, name, email)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Contact mis à jour avec succès !")
}

// ------------------------ 4. Delete Contact ------------------------
func handleDeleteContact(reader *bufio.Reader, store storage.Storer) {
	contacts, _ := store.GetAll()
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à supprimer")
		return
	}

	listContacts(store)
	fmt.Print("\nID du contact à supprimer: ")
	id, _ := strconv.Atoi(readLine(reader))

	err := store.Delete(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Contact supprimé")
}
