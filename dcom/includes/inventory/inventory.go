package inventory

// START OMIT
type Service interface {
	GetOrder(GetOrderRequest, *GetOrderResponse) error
	CreateOrder(CreateOrderRequest, *CreateOrderResponse) error

	OrderStatus(OrderStatusRequest, *OrderStatusResponse) error
	CancelOrder(CancelOrderRequest, *CancelOrderResponse) error

	GetProduct(GetProductRequest, *GetProductResponse) error
	CreateProduct(CreateProductRequest, *CreateProductResponse) error
	UpdateProduct(UpdateProductRequest, *UpdateProductResponse) error
	DeleteProduct(DeleteProductRequest, *DeleteProductResponse) error

	GetSupplier(GetSupplierRequest, *GetSupplierResponse) error
	CreateSupplier(CreateSupplierRequest, *CreateSupplierResponse) error
	UpdateSupplier(UpdateSupplierRequest, *UpdateSupplierResponse) error
	DeleteSupplier(DeleteSupplierRequest, *DeleteSupplierResponse) error
}

// END OMIT

type GetOrderRequest struct{}
type GetOrderResponse struct{}
type OrderStatusRequest struct{}
type OrderStatusResponse struct{}
type CancelOrderRequest struct{}
type CancelOrderResponse struct{}
type GetProductRequest struct{}
type GetProductResponse struct{}
type CreateProductRequest struct{}
type CreateProductResponse struct{}
type CreateOrderRequest struct{}
type CreateOrderResponse struct{}
type UpdateProductRequest struct{}
type UpdateProductResponse struct{}
type DeleteProductRequest struct{}
type DeleteProductResponse struct{}
type GetSupplierRequest struct{}
type GetSupplierResponse struct{}
type CreateSupplierRequest struct{}
type CreateSupplierResponse struct{}
type UpdateSupplierRequest struct{}
type UpdateSupplierResponse struct{}
type DeleteSupplierRequest struct{}
type DeleteSupplierResponse struct{}
