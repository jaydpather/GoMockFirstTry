package employeeService

import(
	"fmt"
	"testing"
	"github.com/golang/mock/gomock"
	"GoMockFirstTry/src/mock_employeeRepository"
)

func formatErrMsg_WrongReturnValue(expected int, actual int) string {
	return fmt.Sprintf("expected %d, but got %d", expected, actual)
}

func TestShouldReturnZeroIfWorking(t *testing.T) {
	actual := ShouldReturnZeroIfWorking()
	expected := 0

	if(expected != actual){
		t.Error(formatErrMsg_WrongReturnValue(expected, actual))
	}
}

func TestInsertEmployee(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish() //this calls our asserts at the end of the function

	mockRepo := mock_employeeRepository.NewMockIEmployeeRepository(mockController)

	expectedValue:= 1
	mockRepo.EXPECT().InsertEmployee().Return(expectedValue)

	actualValue := InsertEmployee(mockRepo)
	if(expectedValue != actualValue) {
		t.Error(formatErrMsg_WrongReturnValue(expectedValue, actualValue))
	}

}


