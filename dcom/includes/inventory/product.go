package inventory

// START OMIT
// A Product is an item that our company stocks
type Product struct {
	ID      int
	Name    string
	SKU     string
	LotSize int
}

// ProductStorage defines the behaviors required to
// perform CRUD operations on an Order
type ProductStorage interface {
	Get(id int) (*Product, error)
	Create(p Product) (*Product, error)
	Update(p *Product) error
	Delete(p *Product) error
}

// END OMIT
