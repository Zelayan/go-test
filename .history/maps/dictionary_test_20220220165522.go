package maps

import (
	"testing"

	"github.com/hashicorp/consul/command/intention/check"
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
		check.Error(err, ErrNotFound)
    })
}

func TestAdd(t *testing.T) {
	check := assert.New(t)

	dictionary := Dictionary{}
    dictionary.Add("test", "this is just a test")
	s, err := dictionary.Search("test")
	
	check.Nil(err)
	check.eq
}
