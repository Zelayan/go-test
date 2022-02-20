package maps

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string err {
    return d[word], nil
}
