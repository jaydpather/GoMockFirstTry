package employeeRepository


type IEmployeeRepository interface {
	InsertEmployee() int
}

type EmployeeRepository struct {
}

func (employeeRepo EmployeeRepository) InsertEmployee() int {
	
	return 1
}