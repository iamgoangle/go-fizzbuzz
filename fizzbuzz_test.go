package fizzbuzz

import (
	"testing"
)

func TestFizzBuzzShouldSayOne(t *testing.T) {
	got := Say(1)
	want := "1"

	if got != want {
		t.Errorf("it should say %q but got %q", got, want)
	}
}

func TestFizzBuzzShouldSayTwo(t *testing.T) {
	got := Say(2)
	want := "2"

	if got != want {
		t.Errorf("it should say %q but got %q", got, want)
	}
}

func TestFizzBuzzShouldSayFizz(t *testing.T) {
	got := Say(3)
	want := "Fizz"

	if got != want {
		t.Errorf("it should say %q but got %q", got, want)
	}
}

func TestFizzBuzzShouldSayBuzz(t *testing.T) {
	got := Say(5)
	want := "Buzz"

	if got != want {
		t.Errorf("it should say %q but got %q", got, want)
	}
}
