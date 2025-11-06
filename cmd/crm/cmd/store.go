package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/vcircosta/GO-TP1/internal/config"
	"github.com/vcircosta/GO-TP1/internal/storage"
)

func initStore() storage.Storer {
	// charge la config si présente (ne fait pas fail hard)
	var cfg *config.Config
	if c, err := config.LoadConfig(); err == nil {
		cfg = c
	} else {
		if verbose {
			log.Printf("Aucun fichier de config chargé: %v", err)
		}
	}

	// priorité : flag -> env -> config -> défaut
	storageType := "json"
	if f := rootCmd.PersistentFlags().Lookup("storage"); f != nil && f.Value.String() != "" {
		storageType = f.Value.String()
	} else if v := os.Getenv("CRM_STORAGE_TYPE"); v != "" {
		storageType = v
	} else if cfg != nil && cfg.Storage.Type != "" {
		storageType = cfg.Storage.Type
	}

	dir := "data"
	if f := rootCmd.PersistentFlags().Lookup("data-dir"); f != nil && f.Value.String() != "" {
		dir = f.Value.String()
	} else if v := os.Getenv("CRM_DATA_DIR"); v != "" {
		dir = v
	} else if cfg != nil && cfg.Storage.DataDir != "" {
		dir = cfg.Storage.DataDir
	}

	sqliteFile := "contacts.db"
	if f := rootCmd.PersistentFlags().Lookup("sqlite-file"); f != nil && f.Value.String() != "" {
		sqliteFile = f.Value.String()
	} else if v := os.Getenv("CRM_SQLITE_FILE"); v != "" {
		sqliteFile = v
	} else if cfg != nil && cfg.Storage.SQLiteFile != "" {
		sqliteFile = cfg.Storage.SQLiteFile
	}

	verboseFlag := false
	if f := rootCmd.PersistentFlags().Lookup("verbose"); f != nil && f.Value.String() == "true" {
		verboseFlag = true
	}

	// assure le répertoire
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Erreur lors de la création du dossier %s: %v", dir, err)
	}

	switch storageType {
	case "gorm":
		dbPath := filepath.Join(dir, sqliteFile)
		store, err := storage.NewGormStore(dbPath)
		if err != nil {
			log.Fatalf("Erreur lors de l'initialisation du store GORM: %v", err)
		}
		if verboseFlag {
			log.Printf("Utilisation du fichier SQLite: %s\n", dbPath)
		}
		return store

	case "json":
		jsonFile := filepath.Join(dir, "contacts.json")
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
