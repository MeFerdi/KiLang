package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.ReadFile("./tests/file.ksm")
	code := string(file)
	fmt.Printf("cd: %s\n", code)
}

