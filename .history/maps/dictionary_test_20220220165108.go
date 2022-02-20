package maps

import (
	"testing"


	"github.com/stretchr/testify/assert"

)

func TestSearch(t *testing.T) {

	check := assert.New(t)
	dictionary := Dictionary{"test":"this is just a test"}
	t.Run("known word", func(t *testing.T) {
        got, _ := dictionary.Search("test")
        want := "this is just a test"
		check.Equal(want, got)
    })

	t.Run("unknown word", func(t *testing.T) {
        _, err := dictionary.Search("unknown")
        want := "could not find the word you were looking for"

		check.Error()
    })


}
