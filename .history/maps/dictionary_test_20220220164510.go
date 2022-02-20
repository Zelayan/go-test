package maps

import (
	"testing"


	"github.com/stretchr/testify/assert"

)

func TestSearch(t *testing.T) {

	check := assert.New(t)
	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")

	want := "this is just a test"

	check.Equal()

	
}
