package postgres

import (
	"database/sql"
	"inventory"
)

// Compile-time proof of interface implementation
var _ inventory.SupplierStorage = (*SupplierService)(nil)

type SupplierService struct {
	db *sql.DB
}

func NewSupplierService(db *sql.DB) inventory.SupplierStorage {
	return &SupplierService{db: db}
}
func (svc *SupplierService) Get(id int) (*inventory.Supplier, error) {
	panic("not implemented")
}

func (svc *SupplierService) Create(s inventory.Supplier) (*inventory.Supplier, error) {
	panic("not implemented")
}

func (svc *SupplierService) Update(s *inventory.Supplier) error {
	panic("not implemented")
}

func (svc *SupplierService) Delete(s *inventory.Supplier) error {
	panic("not implemented")
}
