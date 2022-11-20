package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowerReader struct {
	contents string
	pos      int
}

//extending the ReadAll interface to read one character per line
//here Read function is from io.ReadAll implementation
func (m *MySlowerReader) Read(p []byte) (n int, err error) {
	if m.pos+1 <= len(m.contents) {
		n := copy(p, m.contents[m.pos:m.pos+1]) // copy takes 2 args (dest, src)
		m.pos++
		return n, nil
	}
	return 0, io.EOF
}

func main() {
	MySlowerReaderInstance := &MySlowerReader{
		contents: "Hello World",
	}

	out, err := io.ReadAll(MySlowerReaderInstance)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
