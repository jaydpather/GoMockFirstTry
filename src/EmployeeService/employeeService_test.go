package employeeService

import(
	"fmt"
	"testing"
	//"gomock"
	//"mock_employeeRepository"
)

func TestShouldReturnZeroIfWorking(t *testing.T) {
	actual := ShouldReturnZeroIfWorking()
	expected := 0

	if(actual != expected){
		t.Error(fmt.Sprintf("expected %d, but got %d", expected, actual))
	}
	
}

// func TestInsertEmployee(t *testing.T) {
// 	mockController = gomock.NewController(t)
// 	defer mockController.Finish() //this calls our asserts at the end of the function

// 	mockRepo := NewMockIEmployeeRepository(mockController)

// 	mockRepo.EXPECT().InsertEmployee().Return(1)
// }


