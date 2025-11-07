package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/vcircosta/GO-TP1/internal/storage"
)

func initStore() storage.Storer {
	storageType := "json"

	const configFile = "config_storage.txt"

	if data, err := os.ReadFile(configFile); err == nil {
		storageType = strings.TrimSpace(string(data))
	}

	if flag := rootCmd.Flag("storage"); flag != nil && flag.Changed {
		storageType = flag.Value.String()

		if err := os.WriteFile(configFile, []byte(storageType), 0644); err != nil {
			log.Fatalf("Impossible de sauvegarder le type de stockage : %v", err)
		}
	}

	dir := "data"
	if f := rootCmd.PersistentFlags().Lookup("data-dir"); f != nil && f.Value.String() != "" {
		dir = f.Value.String()
	} else if v := os.Getenv("CRM_DATA_DIR"); v != "" {
		dir = v
	}

	fileName := "contacts"
	if f := rootCmd.PersistentFlags().Lookup("file-name"); f != nil && f.Value.String() != "" {
		fileName = f.Value.String()
	} else if v := os.Getenv("CRM_FILE_NAME"); v != "" {
		fileName = v
	}

	verboseFlag := false
	if f := rootCmd.PersistentFlags().Lookup("verbose"); f != nil && f.Value.String() == "true" {
		verboseFlag = true
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Erreur lors de la création du dossier %s: %v", dir, err)
	}

	switch storageType {
	case "gorm":
		dbPath := filepath.Join(dir, fileName+".db")
		store, err := storage.NewGormStore(dbPath)
		if err != nil {
			log.Fatalf("Erreur lors de l'initialisation du store GORM: %v", err)
		}
		if verboseFlag {
			log.Printf("Utilisation du fichier SQLite: %s\n", dbPath)
		}
		return store

	case "json":
		jsonFile := filepath.Join(dir, fileName+".json")
		store, err := storage.NewJSONStore(jsonFile)
		if err != nil {
			log.Fatalf("Erreur lors de l'initialisation du store JSON: %v", err)
		}
		if verboseFlag {
			log.Printf("Utilisation du fichier JSON: %s\n", jsonFile)
		}
		return store

	default:
		if verboseFlag {
			log.Println("Utilisation du stockage en mémoire")
		}
		return storage.NewMemoryStore()
	}
}
