package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/vcircosta/GO-TP1/internal/app"
	"github.com/vcircosta/GO-TP1/internal/storage"
)

var addCmd = &cobra.Command{
	Use:   "add [nom] [email]",
	Short: "Ajoute un nouveau contact",
	Long: `Ajoute un nouveau contact avec le nom et l'email spécifiés.
	
Exemples :
  crm add "John Doe" "john@example.com" # Ajoute directement un cotact avec le nom "John Doe" et l'email "john@example.com"
  crm add                				# Lance le mode interactif`,
	Run: func(cmd *cobra.Command, args []string) {
		store := initStore()

		switch len(args) {
		case 2:
			// --- Cas correct : ajout direct ---
			name := args[0]
			email := args[1]

			if err := store.Add(&storage.Contact{Name: name, Email: email}); err != nil {
				log.Fatalf("Erreur lors de l'ajout du contact: %v", err)
			}

			fmt.Println("Contact ajouté avec succès !")

			if verbose {
				contacts, _ := store.GetAll()
				fmt.Printf("Total des contacts : %d\n", len(contacts))
			}
		case 0:
			// --- Cas sans argument : mode interactif ---
			reader := bufio.NewReader(os.Stdin)
			app.HandleNewContact(reader, store)

		default:
			// --- Cas d'erreur : mauvais nombre d'arguments ---
			fmt.Println("Nombre de paramètres invalide !!")
			fmt.Println("Utilisation correcte : crm add <nom> <email>")
			fmt.Println("Ou bien : crm add (pour le mode interactif)\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
