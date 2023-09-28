package integers
import "testing"

func TestAdder(t *testing.T){
	var sum int
	sum = Add(2,2)
	expected := 4

	if sum != expected{
		t.Errorf("Expected '%q' but got '%q' ",expected,sum)
	}
}