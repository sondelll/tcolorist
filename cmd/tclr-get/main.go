package main

import (
	"fmt"
	"os"

	"github.com/sondelll/tcolorist/private/libtcolorist"
)

func main() {
	vault := libtcolorist.GetVault()
	err := vault.GetColors()
	if err != nil {
		panic(err)
	}
	os.Exit((0))

	argc := len(os.Args)
	input := os.Args[argc-1]
	clr, err := libtcolorist.ParseColor8Bit(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v,%v,%v\n", clr.Red, clr.Green, clr.Blue)
}
