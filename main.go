package main

import (
	"bufio"
	"flag"
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

var contacts = make(map[int]Contact)

var reader = bufio.NewReader(os.Stdin)

func main() {
	idFlag := flag.Int("id", 0, "ID du contact")
	nameFlag := flag.String("name", "", "Nom du contact")
	emailFlag := flag.String("email", "", "Email du contact")
	flag.Parse()

	if *idFlag != 0 && *nameFlag != "" && *emailFlag != "" {
		addContact(*idFlag, *nameFlag, *emailFlag)
	}

	for {
		fmt.Println("\n--- Mini CRM ---")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister tous les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre à jour un contact")
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
			handleAddContact()
		case "2":
			listContacts()
		case "3":
			handleDeleteContact()
		case "4":
			handleUpdateContact()
		case "5":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

func handleAddContact() {
	fmt.Print("ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	if _, ok := contacts[id]; ok {
		fmt.Println("Un contact avec cet ID existe déjà")
		return
	}

	fmt.Print("Nom: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')

	addContact(id, strings.TrimSpace(name), strings.TrimSpace(email))
}

func addContact(id int, name, email string) {
	contacts[id] = Contact{ID: id, Name: name, Email: email}
	fmt.Println("Contact ajouté avec succès !")
}

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

func handleDeleteContact() {
	fmt.Print("ID du contact à supprimer: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	if _, ok := contacts[id]; ok {
		delete(contacts, id)
		fmt.Println("Contact supprimé")
	} else {
		fmt.Println("Aucun contact avec cet ID")
	}
}

func handleUpdateContact() {
	fmt.Print("ID du contact à mettre à jour: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	if c, ok := contacts[id]; ok {
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

		contacts[id] = Contact{ID: id, Name: name, Email: email}
		fmt.Println("Contact mis à jour !")
	} else {
		fmt.Println("Aucun contact avec cet ID")
	}
}
