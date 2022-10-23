package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func findNumber(n *big.Int, r *big.Int) {
	start := time.Now()
	for c := n; c.Cmp(r) != 0; {
		c = c.Sub(n, big.NewInt(1))
	}
	elapsed := time.Since(start)
	fmt.Print("time = ")
	fmt.Println(elapsed)
}

func main() {
	for keyLength := 8; keyLength <= 4096; keyLength *= 2 {
		i, e := big.NewInt(2), big.NewInt(int64(keyLength))
		amount := i.Exp(i, e, nil)
		fmt.Printf("keyLength = %d \namount = %d \n\n", keyLength, amount)

	}
	for keyLength := 8; keyLength <= 4096; keyLength *= 2 {
		number := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(keyLength)), nil)
		randNumber := new(big.Int)
		randNumber = randNumber.Rand(rand.New(rand.NewSource(21)), number)
		fmt.Printf("randNumber = %X\n", randNumber)
		findNumber(number, randNumber)
	}
}
