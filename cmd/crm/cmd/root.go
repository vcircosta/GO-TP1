package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/vcircosta/GO-TP1/internal/app"
	"github.com/vcircosta/GO-TP1/internal/storage"
)

var (
	dataDir     string
	storageType string
	fileName    string
	verbose     bool

	rootCmd = &cobra.Command{
		Use:   "crm",
		Short: "Mini CRM - Gestionnaire de contacts",
		Long: `Mini CRM est un gestionnaire de contacts en ligne de commande.
		Il permet d'ajouter, lister, mettre à jour et supprimer des contacts.
				
		Exemple d'utilisation:
		crm                             # Lance le mode interactif
		crm add <name> <email>          # Ajoute un contact
		crm list                        # Liste tous les contacts
		crm update <id> <name> <email>  # Met à jour un contact
		crm delete <id>                 # Supprime un contact
		crm set-storage <type>          # Change le stockage par défaut`,
		Run: func(cmd *cobra.Command, args []string) {
			// Choix du store via flag --storage (ou valeur par défaut "json")
			log.Printf("Type de stockage choisi: %s", storageType)

			if err := os.MkdirAll(dataDir, 0o755); err != nil {
				log.Fatalf("Erreur lors de la création du dossier %s: %v", dataDir, err)
			}

			var store storage.Storer
			var err error

			switch storageType {
			case "gorm":
				dbPath := filepath.Join(dataDir, fileName)
				store, err = storage.NewGormStore(dbPath)
				if err != nil {
					log.Fatalf("Erreur init GORM store: %v", err)
				}
				log.Printf("Utilisation GORM avec SQLite: %s", dbPath)
			case "json":
				jsonFile := filepath.Join(dataDir, "contacts.json")
				store, err = storage.NewJSONStore(jsonFile)
				if err != nil {
					log.Fatalf("Erreur init JSON store: %v", err)
				}
				log.Printf("Utilisation JSON: %s", jsonFile)
			default:
				log.Printf("Type de stockage inconnu (%s), utilisation mémoire", storageType)
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
	rootCmd.PersistentFlags().StringVar(&storageType, "storage", "json", "type de stockage: json|gorm|memory")
	rootCmd.PersistentFlags().StringVar(&dataDir, "data-dir", "data", "répertoire des données (utilisé si --storage=gorm ou --storage=json)")
	rootCmd.PersistentFlags().StringVar(&fileName, "file-name", "contacts", "nom du fichier sans l'extension (utilisé si --storage=gorm ou --storage=json)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "mode verbeux")
}
