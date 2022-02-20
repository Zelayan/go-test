package maps

import (
	"testing"


	"github.com/stretchr/testify/assert"

)

func TestSearch(t *testing.T) {

	
	dictionary := Dictionary{"test":"this is just a test"}
	t.Run("known word", func(t *testing.T) {
        got, _ := dictionary.Search("test")
        want := "this is just a test"

		check := assert.New(t)
		check.Equal(want, got)
    })
	
}
