package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste tous les contacts",
	Run: func(cmd *cobra.Command, args []string) {
		store := initStore()
		contacts, err := store.GetAll()
		if err != nil {
			log.Fatalf("Erreur lors de la récupération des contacts: %v", err)
		}
		for _, c := range contacts {
			fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Name, c.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
