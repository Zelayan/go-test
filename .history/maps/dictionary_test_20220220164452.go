package maps

import (
	"testing"

	"github.com/crossdock/crossdock-go/assert"
	"gopkg.in/check.v1"
)

func TestSearch(t *testing.T) {

	check := assert.
	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")

	want := "this is just a test"

	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}

	
}
