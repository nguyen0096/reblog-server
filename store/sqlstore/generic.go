package sqlstore

import "gorm.io/gorm"

type GenericSqlStore interface {
	Create()
	// Read()
	// Update()
	// Delete()
}

type genericSqlStore struct {
	db *gorm.DB
}

func NewGenericSqlStore(db *gorm.DB) GenericSqlStore {
	return &genericSqlStore{
		db: db,
	}
}

func (c *genericSqlStore) Create() {

}
