package exampleService

import (
	"../employeeRepository"
)

func ShouldReturnZeroIfWorking() int {
	return 0
}

func InsertEmployee(repo employeeRepository.IEmployeeRepository) int {
	return repo.InsertEmployee()
}