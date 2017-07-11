package http

import (
	"inventory"
	"inventory/transport"
	"net"
	nh "net/http"
)

// START OMIT

// Compile-time proof of interface implementation
var _ transport.InventoryTransporter = (*RESTService)(nil)

// END OMIT

type RESTService struct {
	orderStore      inventory.OrderStorage
	productStore    inventory.ProductStorage
	supplierStore   inventory.SupplierStorage
	supplierService inventory.SupplierService
}

func NewRESTService(orderStore inventory.OrderStorage, supplierStore inventory.SupplierStorage, supplierService inventory.SupplierService, productStore inventory.ProductStorage) *RESTService {
	return &RESTService{
		orderStore:      orderStore,
		productStore:    productStore,
		supplierStore:   supplierStore,
		supplierService: supplierService,
	}

}

func (svc *RESTService) Serve(l net.Listener) error {
	server := &nh.Server{}
	nh.Handle("/api", svc)
	server.Serve(l)
	return nil
}

func (svc *RESTService) ServeHTTP(res nh.ResponseWriter, req *nh.Request) {
	res.Write([]byte("not implemented"))
}

func (svc *RESTService) GetOrder(inventory.GetOrderRequest, *inventory.GetOrderResponse) error {
	panic("not implemented")
}

func (svc *RESTService) CreateOrder(inventory.CreateOrderRequest, *inventory.CreateOrderResponse) error {
	panic("not implemented")
}

func (svc *RESTService) OrderStatus(inventory.OrderStatusRequest, *inventory.OrderStatusResponse) error {
	panic("not implemented")
}

func (svc *RESTService) CancelOrder(inventory.CancelOrderRequest, *inventory.CancelOrderResponse) error {
	panic("not implemented")
}

func (svc *RESTService) GetProduct(inventory.GetProductRequest, *inventory.GetProductResponse) error {
	panic("not implemented")
}

func (svc *RESTService) CreateProduct(inventory.CreateProductRequest, *inventory.CreateProductResponse) error {
	panic("not implemented")
}

func (svc *RESTService) UpdateProduct(inventory.UpdateProductRequest, *inventory.UpdateProductResponse) error {
	panic("not implemented")
}

func (svc *RESTService) DeleteProduct(inventory.DeleteProductRequest, *inventory.DeleteProductResponse) error {
	panic("not implemented")
}

func (svc *RESTService) GetSupplier(inventory.GetSupplierRequest, *inventory.GetSupplierResponse) error {
	panic("not implemented")
}

func (svc *RESTService) CreateSupplier(inventory.CreateSupplierRequest, *inventory.CreateSupplierResponse) error {
	panic("not implemented")
}

func (svc *RESTService) UpdateSupplier(inventory.UpdateSupplierRequest, *inventory.UpdateSupplierResponse) error {
	panic("not implemented")
}

func (svc *RESTService) DeleteSupplier(inventory.DeleteSupplierRequest, *inventory.DeleteSupplierResponse) error {
	panic("not implemented")
}
