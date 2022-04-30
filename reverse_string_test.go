package reverse_string_test

import (
	"testing"

	reverse_string "github.com/latugovskaya-anastasiya/golang-united-school-homework-4.1"
)

func TestReverseString(t *testing.T) {
	got := reverse_string.ReverseString("Hello")
	want := "olleH"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
