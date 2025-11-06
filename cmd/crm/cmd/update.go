package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <id> <nouveau-nom> <nouveau-email>",
	Short: "Met à jour un contact existant",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("ID invalide: %v", err)
		}
		store := initStore()
		if err := store.Update(id, args[1], args[2]); err != nil {
			log.Fatalf("Erreur lors de la mise à jour: %v", err)
		}
		if verbose {
			fmt.Printf("Contact %d mis à jour avec succès\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
