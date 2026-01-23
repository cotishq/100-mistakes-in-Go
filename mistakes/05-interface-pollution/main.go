package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	source := strings.NewReader("hello")
	dest := bytes.NewBuffer(nil)

	err := copySourceToDest(source, dest)
	if err != nil {
		panic(err)
	}

	fmt.Println("copied output", dest.String())
}

func copySourceToDest(source io.Reader, dest io.Writer) error {
	_, err := io.Copy(dest, source)
	return err
 }
