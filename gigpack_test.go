package gigapack

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	want := 3
	if got := Add(1, 2); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}
