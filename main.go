package main

import (
	"fmt"
	"yana-golang/handler"
	"yana-golang/repository"
	"yana-golang/service"
)

func main() {
	customerMockData := repository.InitCustomerMock()
	customerService := service.InitCustomerService(customerMockData)

	// handler call service
	customers, _ := customerService.GetAllCustomer()
	for i, cus := range customers {
		fmt.Printf("Customers[%v]: %v\n", i, cus.CustomerID)
	}
	err := customerService.CreateCustomer(service.CustomerRequest{Name: "wealthy", Age: 35})
	if err != nil {
		handler.HandleError(err)
	}

	_, err = customerService.GetCustomer("220")
	// handler::Header Response
	if err != nil {
		handler.HandleError(err)
	}

}
