package service

type CustomerRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type CustomerResponse struct {
	CustomerID string `json:"customer_id"`
}

type CustomerService interface {
	CreateCustomer(customer CustomerRequest) error
	GetAllCustomer() ([]CustomerResponse, error)
	GetCustomer(customerId string) (*CustomerResponse, error)
}
