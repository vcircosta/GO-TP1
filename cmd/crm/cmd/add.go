package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/vcircosta/GO-TP1/internal/storage"
)

var addCmd = &cobra.Command{
	Use:   "add <nom> <email>",
	Short: "Ajoute un nouveau contact",
	Long: `Ajoute un nouveau contact avec le nom et l'email spécifiés.
	
Exemple: crm add "John Doe" "john@example.com"`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		store := initStore()
		contact := &storage.Contact{
			Name:  args[0],
			Email: args[1],
		}
		if err := store.Add(contact); err != nil {
			log.Fatalf("Erreur lors de l'ajout du contact: %v", err)
		}
		if verbose {
			fmt.Printf("Contact ajouté avec succès - ID: %d\n", contact.ID)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
