package maps

import (
	"testing"


	"github.com/stretchr/testify/assert"

)

func TestSearch(t *testing.T) {

	check := assert.New(t)
	dictionary := Dictionary{"test", "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		check(want, got)
	}()})

	
}
