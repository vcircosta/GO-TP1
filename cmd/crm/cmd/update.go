package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vcircosta/GO-TP1/internal/app"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] [nouveau-nom] [nouvel-email]",
	Short: "Met à jour un contact existant",
	Long: `Met à jour un contact à partir de son ID.

Exemples :
  crm update 1 "John Smith" "john.smith@example.com" # Modifie le cotact d'ID 1 : le nom par "John Doe" et l'email par "john@example.com"
  crm update                						 # Lance le mode interactif`,
	Run: func(cmd *cobra.Command, args []string) {
		store := initStore()

		switch len(args) {
		case 3:
			// --- Cas correct : modification direct ---
			id, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("ID invalide : %v", err)
			}

			if err := store.Update(id, args[1], args[2]); err != nil {
				log.Fatalf("Erreur lors de la mise à jour : %v", err)
			}

			fmt.Printf("Contact %d mis à jour avec succès\n", id)

			if verbose {
				contacts, _ := store.GetAll()
				fmt.Printf("Total des contacts : %d\n", len(contacts))
			}
		case 0:
			// --- Cas sans argument : mode interactif ---
			reader := bufio.NewReader(os.Stdin)
			app.HandleUpdateContact(reader, store)

		default:
			// --- Cas d'erreur : mauvais nombre d'arguments ---
			fmt.Println("Nombre de paramètres invalide.")
			fmt.Println("Utilisation correcte : crm update <id> <nom> <email>")
			fmt.Println("Ou bien : crm update (pour le mode interactif)\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
