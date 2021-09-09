package main

import "fmt"

var one = "Uno"

func main() {

	fmt.Println(one)

	myFunc()
}

func myFunc() {
	fmt.Println(one)
}
