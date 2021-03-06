package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := strings.Repeat("a", 10)

	if repeated != expected {
		t.Errorf("expected %s but got %s", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeatedString := Repeat("a", 6)
	fmt.Println(repeatedString)
	//Output:
	// aaaaaa
}
