// Author: Vidhi Panda <vidhimpanda@gmail.com>
// Description :Constants in Go

package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println(s)
	// s = "Hello"
	// Error:  cannot assign to s
	// Once a value is assigned to a constant it cannot be changed.
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
