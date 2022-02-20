package maps

want := "this is just a test"

    if got != want {
        t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
    }

func Search(dictionary map[string]string, word string) string {

	return dictionary[word]

}
