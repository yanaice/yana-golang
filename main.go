package main

import (
	"encoding/json"
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
	// fmt.Printf("%v\n", customers)
	// fmt.Printf("%+v\n", customers)
	bytes, _ := json.Marshal(customers)
	fmt.Println(string(bytes))

	var cusJson interface{}
	json.Unmarshal(bytes, &cusJson)
	fmt.Println(cusJson)

	err := customerService.CreateCustomer(service.CustomerRequest{Name: "wealthy", Age: 35})
	if err != nil {
		handler.HandleError(err)
	}

	// struct => json
	customer, _ := customerService.GetCustomer("20")
	cbytes, _ := json.Marshal(customer)
	fmt.Println("Customer:", string(cbytes))

}
