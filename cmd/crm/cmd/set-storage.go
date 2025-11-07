package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var setStorageCmd = &cobra.Command{
	Use:   "set-storage [type]",
	Short: "Change le type de stockage par défaut",
	Long: `Change le type de stockage par défaut.

Exemples :
  crm set-storage gorm   # Change directement le stockage par défaut en GORM/SQLite
  crm set-storage         # Lance le mode interactif pour choisir entre JSON, GORM ou mémoire`,
	Run: func(cmd *cobra.Command, args []string) {
		const configFile = "config_storage.txt"

		validTypes := []string{"json", "gorm", "memory"}

		switch len(args) {
		case 1:
			// --- Cas correct : un paramètre fourni ---
			newStorage := strings.ToLower(args[0])
			if !contains(validTypes, newStorage) {
				fmt.Printf("Type de stockage invalide : %s\n", newStorage)
				fmt.Println("Types valides : json, gorm, memory")
				return
			}

			if err := os.WriteFile(configFile, []byte(newStorage), 0644); err != nil {
				log.Fatalf("Impossible de sauvegarder le type de stockage : %v", err)
			}

			fmt.Printf("Type de stockage par défaut changé en : %s\n", newStorage)

		case 0:
			// --- Cas interactif ---
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Choisissez le type de stockage par défaut :")
			fmt.Println("1: json")
			fmt.Println("2: gorm")
			fmt.Println("3: memory")
			fmt.Print("Votre choix : ")

			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Erreur lecture entrée : %v", err)
			}
			input = strings.TrimSpace(input)

			var choice string
			switch input {
			case "1":
				choice = "json"
			case "2":
				choice = "gorm"
			case "3":
				choice = "memory"
			default:
				fmt.Println("Choix invalide ! Veuillez choisir 1, 2 ou 3.")
				return
			}

			if err := os.WriteFile(configFile, []byte(choice), 0644); err != nil {
				log.Fatalf("Impossible de sauvegarder le type de stockage : %v", err)
			}
			fmt.Printf("Type de stockage par défaut changé en : %s\n", choice)

		default:
			// --- Cas d'erreur : trop d'arguments ---
			fmt.Println("Nombre de paramètres invalide !")
			fmt.Println("Utilisation correcte : crm set-storage <json|gorm|memory>")
			fmt.Println("Ou bien : crm set-storage (pour le mode interactif)\n")
		}
	},
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(setStorageCmd)
}
