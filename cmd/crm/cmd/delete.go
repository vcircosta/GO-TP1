package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Supprime un contact",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("ID invalide: %v", err)
		}
		store := initStore()
		if err := store.Delete(id); err != nil {
			log.Fatalf("Erreur lors de la suppression: %v", err)
		}
		if verbose {
			fmt.Printf("Contact %d supprimé avec succès\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
