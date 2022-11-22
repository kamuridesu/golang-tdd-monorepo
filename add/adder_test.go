package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Testing 2+2", func(t *testing.T) {
		want := 4
		got := Add(2, 2)
		if want != got {
			t.Errorf("expected '%d' but got '%d'", want, got)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// output: 6
}
