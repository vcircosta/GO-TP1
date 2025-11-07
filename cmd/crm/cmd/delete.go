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

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Supprime un contact",
	Long: `Supprime un contact à partir de son ID.

Exemples :
  crm delete 1        # Supprime directement le contact d'ID 1
  crm delete          # Lance le mode interactif`,
	Run: func(cmd *cobra.Command, args []string) {
		store := initStore()

		switch len(args) {
		case 1:
			// --- Cas correct : suppression direct ---
			id, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("ID invalide : %v", err)
			}

			if err := store.Delete(id); err != nil {
				log.Fatalf("Erreur lors de la suppression : %v", err)
			}

			fmt.Printf("Contact %d supprimé avec succès\n", id)

			if verbose {
				contacts, _ := store.GetAll()
				fmt.Printf("Total des contacts : %d\n", len(contacts))
			}
		case 0:
			// --- Cas sans argument : mode interactif ---
			reader := bufio.NewReader(os.Stdin)
			app.HandleDeleteContact(reader, store)

		default:
			// --- Cas d'erreur : mauvais nombre d'arguments ---
			fmt.Println("Nombre de paramètres invalide !!")
			fmt.Println("Utilisation correcte : crm delete <id>")
			fmt.Println("Ou bien : crm delete (pour le mode interactif)\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
