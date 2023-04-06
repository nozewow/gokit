package main

import (
	"fmt"
	"math/big"
)

func main() {
	i1 := big.NewInt(100)
	i2 := big.NewInt(50)

	fmt.Println(i1.Cmp(i2))
}
