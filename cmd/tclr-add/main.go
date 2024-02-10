package main

import (
	"fmt"
	"os"

	"github.com/sondelll/tcolorist/private/libtcolorist"
)

func argsCheck() {
	if len(os.Args) < 2 {
		fmt.Printf("- - tColorist Adder - -\n")
		fmt.Printf("No arguments provided\n")
		os.Exit(0)
	}
}

func main() {
	argsCheck()

	argc := len(os.Args)
	input := os.Args[argc-1]
	vault := libtcolorist.GetVault()
	clr, parseErr := libtcolorist.ParseColor8Bit(input)
	if parseErr != nil {
		panic(parseErr)
	}

	addErr := vault.AddColorRGB8(clr)
	if addErr != nil {
		panic(addErr)
	}

	fmt.Printf("Added RBG8 color (%v,%v,%v)\n", clr.Red, clr.Green, clr.Blue)
}
