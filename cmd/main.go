package main

import (
	"fmt"

	test "github.com/abdulhaseeb08/go-project/tests"
)

//entry point into our program
func main() {
	fmt.Println("hello")
	fmt.Println(test.Distance("AHHHDK", "AHHDK"))
	fmt.Println(test.ToRNA("TCGA"))

	clockM := test.New(5, 6)
	fmt.Println(clockM.String())
}
