package storage

import (
	"os"
	"path/filepath"
	"testing"
)

func newTestGormStore(t *testing.T) (*GormStore, string) {
	tmpFile := filepath.Join(os.TempDir(), "test_contacts_"+t.Name()+".db")
	store, err := NewGormStore(tmpFile)
	if err != nil {
		t.Fatalf("Erreur lors de la création du GormStore: %v", err)
	}
	t.Cleanup(func() {
		os.Remove(tmpFile)
	})
	return store, tmpFile
}

func TestGormAddAndGetAll(t *testing.T) {
	store, _ := newTestGormStore(t)

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

	found := map[string]bool{}
	for _, c := range all {
		found[c.Name] = true
	}
	if !found["Alice"] || !found["Bob"] {
		t.Errorf("Les contacts ajoutés ne correspondent pas")
	}
}

func TestGormGetById(t *testing.T) {
	store, _ := newTestGormStore(t)

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

func TestGormUpdate(t *testing.T) {
	store, _ := newTestGormStore(t)

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

func TestGormDelete(t *testing.T) {
	store, _ := newTestGormStore(t)

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

func TestGormDeleteInexistant(t *testing.T) {
	store, _ := newTestGormStore(t)

	err := store.Delete(99)
	if err == nil {
		t.Error("Suppression d'un contact inexistant devrait retourner une erreur")
	}
}

func TestGormPersistence(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "test_gorm_persistence.db")
	defer os.Remove(tmpFile)

	store1, err := NewGormStore(tmpFile)
	if err != nil {
		t.Fatalf("Erreur création store1: %v", err)
	}

	c1 := &Contact{Name: "Frank", Email: "frank@example.com"}
	c2 := &Contact{Name: "Grace", Email: "grace@example.com"}
	_ = store1.Add(c1)
	_ = store1.Add(c2)

	store2, err := NewGormStore(tmpFile)
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

func TestGormUpdatePersistence(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "test_gorm_update.db")
	defer os.Remove(tmpFile)

	store1, _ := NewGormStore(tmpFile)
	c := &Contact{Name: "Henry", Email: "henry@example.com"}
	_ = store1.Add(c)

	_ = store1.Update(1, "Henry Updated", "henry2@example.com")

	store2, _ := NewGormStore(tmpFile)
	updated, err := store2.GetById(1)
	if err != nil {
		t.Fatalf("Erreur GetById après rechargement: %v", err)
	}
	if updated.Name != "Henry Updated" || updated.Email != "henry2@example.com" {
		t.Errorf("Mise à jour non persistée correctement: %+v", updated)
	}
}

func TestGormNextID(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), "test_gorm_nextid.db")
	defer os.Remove(tmpFile)

	store1, _ := NewGormStore(tmpFile)
	_ = store1.Add(&Contact{Name: "Kate", Email: "kate@example.com"})
	_ = store1.Add(&Contact{Name: "Leo", Email: "leo@example.com"})

	store2, _ := NewGormStore(tmpFile)
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
