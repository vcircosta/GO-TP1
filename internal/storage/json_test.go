package storage

import (
	"os"
	"testing"
)

func newTestJSONStore(t *testing.T) (*JSONStore, string) {
	tmpFile := "test_contacts_" + t.Name() + ".json"
	store, err := NewJSONStore(tmpFile)
	if err != nil {
		t.Fatalf("Erreur lors de la création du JSONStore: %v", err)
	}
	t.Cleanup(func() {
		os.Remove(tmpFile)
	})
	return store, tmpFile
}

func TestJSONAddAndGetAll(t *testing.T) {
	store, _ := newTestJSONStore(t)

	c1 := &Contact{Name: "Alice", Email: "alice@example.com"}
	c2 := &Contact{Name: "Bob", Email: "bob@example.com"}

	if err := store.Add(c1); err != nil {
		t.Fatalf("Erreur lors de l'ajout de c1: %v", err)
	}
	if err := store.Add(c2); err != nil {
		t.Fatalf("Erreur lors de l'ajout de c2: %v", err)
	}

	all, _ := store.GetAll()
	if len(all) != 2 {
		t.Fatalf("Nombre de contacts attendu 2, obtenu %d", len(all))
	}

	found := make(map[string]bool)
	for _, c := range all {
		found[c.Name] = true
	}
	if !found["Alice"] || !found["Bob"] {
		t.Errorf("Les contacts ajoutés ne correspondent pas")
	}
}

func TestJSONGetById(t *testing.T) {
	store, _ := newTestJSONStore(t)

	c := &Contact{Name: "Charlie", Email: "charlie@example.com"}
	_ = store.Add(c)

	got, err := store.GetById(1)
	if err != nil {
		t.Fatalf("Erreur GetById: %v", err)
	}
	if got.Name != "Charlie" {
		t.Errorf("Nom attendu Charlie, obtenu %s", got.Name)
	}
}

func TestJSONUpdate(t *testing.T) {
	store, _ := newTestJSONStore(t)

	c := &Contact{Name: "David", Email: "david@example.com"}
	_ = store.Add(c)

	err := store.Update(1, "David Updated", "david2@example.com")
	if err != nil {
		t.Fatalf("Erreur Update: %v", err)
	}

	got, _ := store.GetById(1)
	if got.Name != "David Updated" || got.Email != "david2@example.com" {
		t.Errorf("Update incorrect: %+v", got)
	}
}

func TestJSONDelete(t *testing.T) {
	store, _ := newTestJSONStore(t)

	c := &Contact{Name: "Eve", Email: "eve@example.com"}
	_ = store.Add(c)

	err := store.Delete(1)
	if err != nil {
		t.Fatalf("Erreur Delete: %v", err)
	}

	_, err = store.GetById(1)
	if err == nil {
		t.Error("Le contact aurait dû être supprimé, mais est encore présent")
	}
}

func TestJSONDeleteInexistant(t *testing.T) {
	store, _ := newTestJSONStore(t)

	err := store.Delete(99)
	if err == nil {
		t.Error("Suppression d'un contact inexistant devrait retourner une erreur")
	}
}

func TestJSONPersistence(t *testing.T) {
	tmpFile := "test_persistence.json"
	defer os.Remove(tmpFile)

	store1, err := NewJSONStore(tmpFile)
	if err != nil {
		t.Fatalf("Erreur création store1: %v", err)
	}

	c1 := &Contact{Name: "Frank", Email: "frank@example.com"}
	c2 := &Contact{Name: "Grace", Email: "grace@example.com"}
	_ = store1.Add(c1)
	_ = store1.Add(c2)

	store2, err := NewJSONStore(tmpFile)
	if err != nil {
		t.Fatalf("Erreur création store2: %v", err)
	}

	all, _ := store2.GetAll()
	if len(all) != 2 {
		t.Fatalf("Nombre de contacts après rechargement attendu 2, obtenu %d", len(all))
	}

	frank, err := store2.GetById(1)
	if err != nil || frank.Name != "Frank" {
		t.Errorf("Contact Frank non trouvé après rechargement")
	}

	grace, err := store2.GetById(2)
	if err != nil || grace.Name != "Grace" {
		t.Errorf("Contact Grace non trouvé après rechargement")
	}
}

func TestJSONUpdatePersistence(t *testing.T) {
	tmpFile := "test_update_persistence.json"
	defer os.Remove(tmpFile)

	store1, _ := NewJSONStore(tmpFile)
	c := &Contact{Name: "Henry", Email: "henry@example.com"}
	_ = store1.Add(c)

	_ = store1.Update(1, "Henry Updated", "henry2@example.com")

	store2, _ := NewJSONStore(tmpFile)
	updated, err := store2.GetById(1)
	if err != nil {
		t.Fatalf("Erreur GetById après rechargement: %v", err)
	}
	if updated.Name != "Henry Updated" || updated.Email != "henry2@example.com" {
		t.Errorf("Mise à jour non persistée correctement: %+v", updated)
	}
}

func TestJSONDeletePersistence(t *testing.T) {
	tmpFile := "test_delete_persistence.json"
	defer os.Remove(tmpFile)

	store1, _ := NewJSONStore(tmpFile)
	c := &Contact{Name: "Ivy", Email: "ivy@example.com"}
	_ = store1.Add(c)
	_ = store1.Delete(1)

	store2, _ := NewJSONStore(tmpFile)
	_, err := store2.GetById(1)
	if err == nil {
		t.Error("Le contact supprimé est réapparu après rechargement")
	}

	all, _ := store2.GetAll()
	if len(all) != 0 {
		t.Errorf("Nombre de contacts après suppression et rechargement attendu 0, obtenu %d", len(all))
	}
}

func TestJSONFileNotExists(t *testing.T) {
	tmpFile := "non_existent_file.json"
	defer os.Remove(tmpFile)

	store, err := NewJSONStore(tmpFile)
	if err != nil {
		t.Fatalf("La création d'un store avec un fichier inexistant ne devrait pas échouer: %v", err)
	}

	all, _ := store.GetAll()
	if len(all) != 0 {
		t.Errorf("Le store devrait être vide, mais contient %d contacts", len(all))
	}

	c := &Contact{Name: "Jack", Email: "jack@example.com"}
	_ = store.Add(c)

	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		t.Error("Le fichier JSON devrait avoir été créé après l'ajout d'un contact")
	}
}

func TestJSONNextIDPersistence(t *testing.T) {
	tmpFile := "test_nextid.json"
	defer os.Remove(tmpFile)

	store1, _ := NewJSONStore(tmpFile)
	_ = store1.Add(&Contact{Name: "Kate", Email: "kate@example.com"})
	_ = store1.Add(&Contact{Name: "Leo", Email: "leo@example.com"})

	store2, _ := NewJSONStore(tmpFile)
	c3 := &Contact{Name: "Mia", Email: "mia@example.com"}
	_ = store2.Add(c3)

	if c3.ID != 3 {
		t.Errorf("L'ID du nouveau contact devrait être 3, obtenu %d", c3.ID)
	}

	all, _ := store2.GetAll()
	if len(all) != 3 {
		t.Errorf("Nombre de contacts attendu 3, obtenu %d", len(all))
	}
}
