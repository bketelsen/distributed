package acme

import "inventory"

// Compile-time proof of interface implementation
var _ inventory.SupplierService = (*AcmeClientService)(nil)

// START OMIT
type AcmeClientService struct {
	URL string
}

func NewClient(url string) inventory.SupplierService {
	return &AcmeClientService{URL: url}
}

func (a *AcmeClientService) PlaceOrder(o *inventory.Order) error {
	panic("not implemented")
}

func (a *AcmeClientService) GetStatus(o *inventory.Order) error {
	panic("not implemented")
}

// END OMIT
