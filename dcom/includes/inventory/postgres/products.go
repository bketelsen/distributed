package postgres

import (
	"database/sql"
	"inventory"
)

// Compile-time proof of interface implementation
var _ inventory.ProductStorage = (*ProductService)(nil)

type ProductService struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) inventory.ProductStorage {
	return &ProductService{db: db}
}

func (svc *ProductService) Get(id int) (*inventory.Product, error) {
	panic("not implemented")
}

func (svc *ProductService) Create(p inventory.Product) (*inventory.Product, error) {
	panic("not implemented")
}

func (svc *ProductService) Update(p *inventory.Product) error {
	panic("not implemented")
}

func (svc *ProductService) Delete(p *inventory.Product) error {
	panic("not implemented")
}
