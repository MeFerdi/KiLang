package main

import (
	"fmt"
	"ksm/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Halo %s! This is the Kisumu programming language!\n", user.Username)
	fmt.Printf("Type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}