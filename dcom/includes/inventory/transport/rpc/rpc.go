package rpc

import (
	"inventory"
	"inventory/transport"
	"net"
	"net/rpc"
)

// Compile-time proof of interface implementation
var _ transport.InventoryTransporter = (*RPCService)(nil)

type RPCService struct {
	orderStore      inventory.OrderStorage
	productStore    inventory.ProductStorage
	supplierStore   inventory.SupplierStorage
	supplierService inventory.SupplierService
}

func NewRPCService(orderStore inventory.OrderStorage, supplierStore inventory.SupplierStorage, supplierService inventory.SupplierService, productStore inventory.ProductStorage) *RPCService {
	return &RPCService{
		orderStore:      orderStore,
		productStore:    productStore,
		supplierStore:   supplierStore,
		supplierService: supplierService,
	}

}

func (svc *RPCService) Serve(l net.Listener) error {
	err := rpc.Register(svc)
	if err != nil {
		return err
	}
	rpc.Accept(l) // blocks
	return nil
}

func (svc *RPCService) GetOrder(inventory.GetOrderRequest, *inventory.GetOrderResponse) error {
	panic("not implemented")
}

func (svc *RPCService) CreateOrder(inventory.CreateOrderRequest, *inventory.CreateOrderResponse) error {
	panic("not implemented")
}

func (svc *RPCService) OrderStatus(inventory.OrderStatusRequest, *inventory.OrderStatusResponse) error {
	panic("not implemented")
}

func (svc *RPCService) CancelOrder(inventory.CancelOrderRequest, *inventory.CancelOrderResponse) error {
	panic("not implemented")
}

func (svc *RPCService) GetProduct(inventory.GetProductRequest, *inventory.GetProductResponse) error {
	panic("not implemented")
}

func (svc *RPCService) CreateProduct(inventory.CreateProductRequest, *inventory.CreateProductResponse) error {
	panic("not implemented")
}

func (svc *RPCService) UpdateProduct(inventory.UpdateProductRequest, *inventory.UpdateProductResponse) error {
	panic("not implemented")
}

func (svc *RPCService) DeleteProduct(inventory.DeleteProductRequest, *inventory.DeleteProductResponse) error {
	panic("not implemented")
}

func (svc *RPCService) GetSupplier(inventory.GetSupplierRequest, *inventory.GetSupplierResponse) error {
	panic("not implemented")
}

func (svc *RPCService) CreateSupplier(inventory.CreateSupplierRequest, *inventory.CreateSupplierResponse) error {
	panic("not implemented")
}

func (svc *RPCService) UpdateSupplier(inventory.UpdateSupplierRequest, *inventory.UpdateSupplierResponse) error {
	panic("not implemented")
}

func (svc *RPCService) DeleteSupplier(inventory.DeleteSupplierRequest, *inventory.DeleteSupplierResponse) error {
	panic("not implemented")
}
