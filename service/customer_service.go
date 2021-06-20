package service

import (
	"yana-golang/errs"
	"yana-golang/repository"
)

type customerService struct {
	cus repository.CustomerRepository
}

func InitCustomerService(cus repository.CustomerRepository) CustomerService {
	return &customerService{
		cus: cus,
	}
}

func (c *customerService) CreateCustomer(customer CustomerRequest) error {
	err := map[string]interface{}{
		"reqBody": customer,
	}
	return errs.ErrorBadRequest("cannot create customer", err)
}

func (c *customerService) GetAllCustomer() ([]CustomerResponse, error) {
	results, err := c.cus.GetAll()
	if err != nil {
		return nil, errs.ErrorInternalServer()
	}

	var customers []CustomerResponse
	for _, customer := range results {
		customers = append(customers, CustomerResponse{
			CustomerID: customer.CustomerID,
		})
	}

	return customers, nil
}

func (c *customerService) GetCustomer(customerId string) (*CustomerResponse, error) {
	customer, err := c.cus.GetById(customerId)
	if err != nil {
		return nil, errs.ErrorNotFound("not found", map[string]interface{}{
			"customerId": customerId,
		})
	}
	response := CustomerResponse{
		CustomerID: customer.CustomerID,
	}
	return &response, nil
}
