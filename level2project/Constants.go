package main

import (
	"fmt"
	"math"
)

const name = "constant"

func main() {
	fmt.Println(name)
	const n = 50000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
