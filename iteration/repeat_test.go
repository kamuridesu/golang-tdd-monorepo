package iteration

import "testing"
import "fmt"
import "strings"

func assert(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func TestRepeatDefault(t *testing.T) {
	want := strings.Repeat("a", 5)
	got := Repeat("a")
	assert(t, got, want)
}

func TestRepeat6(t *testing.T) {
	want := strings.Repeat("a", 6)
	got := Repeat("a", 6)
	assert(t, got, want)
}

func TestRepeat1(t *testing.T) {
	want := strings.Repeat("a", 1)
	got := Repeat("a", 1)
	assert(t, got, want)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeatDefault() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	// output: aaaaa
}

func ExampleRepeat1() {
	repeated := Repeat("a", 1)
	fmt.Println(repeated)
	// output: a
}
