package storage

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(dbPath string) (*GormStore, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto-migrate crée/met à jour la table contacts selon la struct Contact
	if err := db.AutoMigrate(&Contact{}); err != nil {
		return nil, err
	}
	return &GormStore{db: db}, nil
}

func (s *GormStore) Add(contact *Contact) error {
	return s.db.Create(contact).Error
}

func (s *GormStore) GetAll() ([]*Contact, error) {
	var contacts []*Contact
	if err := s.db.Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (s *GormStore) GetById(id int) (*Contact, error) {
	var c Contact
	if err := s.db.First(&c, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrContactNoFound(id)
		}
		return nil, err
	}
	return &c, nil
}

func (s *GormStore) Update(id int, newName string, newEmail string) error {
	result := s.db.Model(&Contact{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":  newName,
		"email": newEmail,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrContactNoFound(id)
	}
	return nil
}

func (s *GormStore) Delete(id int) error {
	result := s.db.Delete(&Contact{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrContactNoFound(id)
	}
	return nil
}
