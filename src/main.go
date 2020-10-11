package main 

import (
	"fmt"
	"GoMockFirstTry/src/employeeService"
	"GoMockFirstTry/src/employeeRepository"
)

func main(){
	fmt.Println("Hello, World!")

	employeeRepo := employeeRepository.EmployeeRepository { }
	employeeService.InsertEmployee(employeeRepo)

}