package main 

import (
	"fmt"
	"./employeeService"
	"./employeeRepository"
)

func main(){
	fmt.Println("Hello, World!")

	employeeRepo := employeeRepository.EmployeeRepository { }
	employeeService.InsertEmployee(employeeRepo)

}