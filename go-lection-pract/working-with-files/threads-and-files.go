package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(file, os.Stdin)
	os.Exit(1)
}
