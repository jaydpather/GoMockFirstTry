package employeeService

import (
	"GoMockFirstTry/src/employeeRepository"
	"fmt"
)

func ShouldReturnZeroIfWorking() int {
	return 0
}

func InsertEmployee(repo employeeRepository.IEmployeeRepository) int {
	fmt.Println("inserting employee")
	return repo.InsertEmployee()
}