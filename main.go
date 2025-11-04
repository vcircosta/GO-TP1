package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vcircosta/GO-TP1/minicrm"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	minicrm.SetTestContacts()

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

		switch strings.TrimSpace(input) {
		case "1":
			minicrm.HandleNewContact()
		case "2":
			minicrm.ListContacts()
		case "3":
			minicrm.HandleUpdateContact()
		case "4":
			minicrm.HandleDeleteContact()
		case "5":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}
