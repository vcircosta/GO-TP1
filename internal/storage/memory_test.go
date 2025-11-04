package storage

import (
	"testing"
)

// Helper : crée un MemoryStore vide prêt à l’emploi
func newTestStore() *MemoryStore {
	return NewMemoryStore()
}

func TestAddAndGetAll(t *testing.T) {
	store := newTestStore()

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

	if all[0].Name != "Alice" || all[1].Email != "bob@example.com" {
		t.Errorf("Les contacts ajoutés ne correspondent pas")
	}
}

func TestGetById(t *testing.T) {
	store := newTestStore()

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

func TestUpdate(t *testing.T) {
	store := newTestStore()
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

func TestDelete(t *testing.T) {
	store := newTestStore()
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

func TestDeleteInexistant(t *testing.T) {
	store := newTestStore()
	err := store.Delete(99)
	if err == nil {
		t.Error("Suppression d'un contact inexistant devrait retourner une erreur")
	}
}
