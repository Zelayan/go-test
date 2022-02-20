package maps

import (
	"testing"

	"github.com/crossdock/crossdock-go/assert"
	"github.com/stretchr/testify/assert"

)

func TestSearch(t *testing.T) {

	check := assert.New(t)
	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")

	want := "this is just a test"

	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}

	
}
