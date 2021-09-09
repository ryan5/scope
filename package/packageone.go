package packageone

import "fmt"

var PackageVar = "Package Var."

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}
