package main

import (
	"fmt"
	"os"

	"git.sr.ht/~alphatroya/atr-capture/env"
)

func main() {
	err := env.CheckEnvs()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Println("Hello, World!")
}
