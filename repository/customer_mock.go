package repository

import "errors"

// ADAPTER
type customerMock struct { // me
	customers []Customer // need to used
}

func InitCustomerMock() CustomerRepository { // what i can do (interface)
	customers :=
		[]Customer{
			{
				CustomerID: "10",
				Name:       "yanee",
				Age:        33,
			},
			{
				CustomerID: "20",
				Name:       "yana",
				Age:        33,
			},
		}
	return customerMock{customers: customers}
}

func (c customerMock) GetAll() ([]Customer, error) {
	return c.customers, nil
}

func (c customerMock) GetById(string) (*Customer, error) {
	return nil, errors.New("!!!ERROR FROM CLIENT")
}
