package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vcircosta/GO-TP1/internal/app"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste tous les contacts",
	Long: `Affiche la liste de tous les contacts enregistrés.

Exemple :
  crm list         # Liste tous les contacts`,
	Run: func(cmd *cobra.Command, args []string) {
		store := initStore()

		switch len(args) {
		case 0:
			// --- Cas correct : affichage des contacts ---
			app.ListContacts(store)

		default:
			// --- Cas d'erreur : s'il y a des arguments ---
			fmt.Println("Cette commande ne prend aucun paramètre.")
			fmt.Println("Utilisation correcte : crm list\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
