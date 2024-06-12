package iteration

import "testing"

func TestRepeat(t *testing.T) {
	retries := Repeat("a")
	expected := "aaaaa"

	if retries != expected {
		t.Errorf("expected '%s' but got '%s'", expected, retries)
	}
}

const RetriesQuantity = 5

func Repeat(character string) string {
	var retries string

	for i := 0; i < RetriesQuantity; i++ {
		retries += character
	}
	return retries
}

func BenchmarkRetries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
