package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/vcircosta/GO-TP1/internal/app"
	"github.com/vcircosta/GO-TP1/internal/storage"
)

func main() {
	var memoryUseJson = true

	if memoryUseJson {
		// JSONStore
		dataDir := "data"
		if err := os.MkdirAll(dataDir, 0755); err != nil {
			log.Fatalf("Erreur lors de la création du dossier %s: %v", dataDir, err)
		}
		jsonFile := filepath.Join(dataDir, "contacts.json")
		store, err := storage.NewJSONStore(jsonFile)
		if err != nil {
			log.Fatalf("Erreur lors de l'initialisation du store JSON: %v", err)
		}
		fmt.Printf("Utilisation du fichier: %s\n", jsonFile)
		app.Run(store)
	} else {
		// MemoryStore
		fmt.Println("Utilisation du stockage en mémoire")
		app.Run(storage.NewMemoryStore())
	}
}
