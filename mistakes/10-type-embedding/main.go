package main 

import (
	"fmt"
	"io"
	"os"
)

// Good: embedding to promote behaviour intentionally

type Logger struct {
	io.WriteCloser
}

func goodEmbedding() {
	logger := Logger{WriteCloser: os.Stdout}

	_, _ = logger.Write([]byte("hello from logger"))
	_ = logger.Close()

	fmt.Println("Logger satisfies io.WriteCloser without boilerplate")
}

func main() {
	goodEmbedding()
}
