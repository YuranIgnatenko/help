package main

import (
	. "help_filego"
	"fmt"
)

var PATH = "modulet.go"

func main() {
	p := fmt.Println
	h := Help{}

	p(h.Dir(PATH, "Cat"))
	p(h.DirVerbose(PATH, "Dog"))

}
