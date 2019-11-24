package integers

import (
	"fmt"
	"testing"
)

// Add takes two integers and returns the sum of them.
func Add(a, b int) int {
	return a + b
}
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected %d but get %d", expected, sum)
	}
}
func ExampleAdd() {
	sum := Add(1, 6)
	fmt.Println(sum)
	//Output:7
}
