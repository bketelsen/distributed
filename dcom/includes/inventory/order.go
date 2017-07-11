package inventory

import "time"

// An Order is a request to purchase a quantity of products
// from a Supplier
type Order struct {
	ID                   int
	Supplier             Supplier
	Products             []Product
	SupplierOrderID      int
	OrderDate            time.Time
	ExpectedDeliveryDate time.Time
}

// OrderStorage defines the behaviors required to
// perform CRUD operations on an Order
type OrderStorage interface {
	Get(id int) (*Order, error)
	Create(o Order) (*Order, error)
	Cancel(o *Order) error
}
