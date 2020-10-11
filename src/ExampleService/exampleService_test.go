package exampleService

import(
	"fmt"
	"testing"
)


func TestShouldReturnZeroIfWorking(t *testing.T) {
	actual := ShouldReturnZeroIfWorking()
	expected := 0

	if(actual != expected){
		t.Error(fmt.Sprintf("expected %d, but got %d", expected, actual))
	}
	
}