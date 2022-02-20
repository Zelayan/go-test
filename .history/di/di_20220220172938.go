package di

func Greet(writer *bytes.Buffer, name string) {
    fmt.Printf("Hello, %s", name)
}