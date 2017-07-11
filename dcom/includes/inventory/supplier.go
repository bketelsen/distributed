package inventory

// A Supplier is a company that provides Products
type Supplier struct {
	ID      int
	Name    string
	Catalog []Product
}

// SupplierStorage defines the behaviors required to
// perform CRUD operations on a Supplier
type SupplierStorage interface {
	Get(id int) (*Supplier, error)
	Create(s Supplier) (*Supplier, error)
	Update(s *Supplier) error
	Delete(s *Supplier) error
}

// SupplierService defines the behaviors required to
// interact with a supplier's order management system
type SupplierService interface {
	PlaceOrder(o *Order) error
	GetStatus(o *Order) error
}
