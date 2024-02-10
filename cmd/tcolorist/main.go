package main

import (
	"fmt"
	"os"

	"github.com/useful-artifacts/tcolorist/private/libtcolorist"
)

func argsCheck() {
	if len(os.Args) < 2 {
		fmt.Printf("- - tColorist - -\n")
		fmt.Printf("No arguments provided\n")
		os.Exit(0)
	}
}

func main() {
	argsCheck()

	argc := len(os.Args)
	input := os.Args[argc-1]
	vault := libtcolorist.GetVault()
	vault.GetColors()
	os.Exit((0))
	clr, err := libtcolorist.ParseColor8Bit(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v,%v,%v\n", clr.Red, clr.Green, clr.Blue)
}
