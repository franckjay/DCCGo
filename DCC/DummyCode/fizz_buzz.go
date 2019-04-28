package main

import (
	"bytes"
	"fmt"
)

// Declare our constant values
const ( fizz int = 3
	buzz int = 5
	fizzbuzz int = 15
)

func main() {
	// Iterate through i to 20, where i:= is an implicit type declaration
	for i := 0; i < 20; i++ {
		var buffer bytes.Buffer
		if i % fizz == 0 {
			buffer.WriteString("fizz")
		}
		if i % buzz == 0 {
			buffer.WriteString("buzz")
		}
		fmt.Printf("For variable %v, we have a %s\n", i, buffer.String())//Write out results
	}

}
