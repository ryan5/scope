package main

import (
	packageone "myapp/package"
)

var myVar = "This is my package level var."

func main() {
	blockVar := "Block var"

	packageone.PrintMe(blockVar, myVar)
}
