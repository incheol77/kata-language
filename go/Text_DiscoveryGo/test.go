package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	b.Write([]byte("good morning"))
	fmt.Fprintf(&b, "go in action")
	io.Copy(os.Stdout, &b)
}

