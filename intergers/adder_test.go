package intergers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(1, 2)
	expected := 3

	if sum != expected {
		t.Errorf("expected '%d' sum '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 4)
	fmt.Println(sum)
	// Output: 6
}
