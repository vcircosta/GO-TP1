package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/vcircosta/GO-TP1/internal/storage"
)

func initStore() storage.Storer {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Fatalf("Erreur lors de la création du dossier %s: %v", dataDir, err)
	}

	if useJSON {
		jsonFile := filepath.Join(dataDir, "contacts.json")
		store, err := storage.NewJSONStore(jsonFile)
		if err != nil {
			log.Fatalf("Erreur lors de l'initialisation du store JSON: %v", err)
		}
		if verbose {
			log.Printf("Utilisation du fichier: %s\n", jsonFile)
		}
		return store
	}

	if verbose {
		log.Println("Utilisation du stockage en mémoire")
	}
	return storage.NewMemoryStore()
}
