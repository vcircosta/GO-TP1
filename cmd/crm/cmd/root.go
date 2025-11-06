package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/vcircosta/GO-TP1/internal/app"
	"github.com/vcircosta/GO-TP1/internal/storage"
)

var (
	dataDir string
	useJSON bool
	verbose bool
	rootCmd = &cobra.Command{
		Use:   "crm",
		Short: "Mini CRM - Gestionnaire de contacts",
		Long: `Mini CRM est un gestionnaire de contacts en ligne de commande.
Il permet d'ajouter, lister, mettre à jour et supprimer des contacts.
		
Exemple d'utilisation:
  crm                    # Lance le mode interactif
  crm add                # Ajoute un contact
  crm list               # Liste tous les contacts
  crm update <id>        # Met à jour un contact
  crm delete <id>        # Supprime un contact`,
		Run: func(cmd *cobra.Command, args []string) {
			// s'assure que dataDir existe
			if err := os.MkdirAll(dataDir, 0o755); err != nil {
				log.Fatalf("Erreur lors de la création du dossier %s: %v", dataDir, err)
			}

			var store storage.Storer
			var err error

			if useJSON {
				jsonFile := filepath.Join(dataDir, "contacts.json")
				store, err = storage.NewJSONStore(jsonFile)
				if err != nil {
					log.Fatalf("Erreur lors de l'initialisation du store JSON: %v", err)
				}
				fmt.Printf("Utilisation du fichier: %s\n\n", jsonFile)
			} else {
				fmt.Println("Utilisation du stockage en mémoire")
				store = storage.NewMemoryStore()
			}

			app.Run(store)
		},
	}
)

// Execute est appelée par main.go
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dataDir, "data-dir", "data", "répertoire des données")
	rootCmd.PersistentFlags().BoolVar(&useJSON, "json", true, "utiliser le store JSON (data/contacts.json)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "mode verbeux")
}
