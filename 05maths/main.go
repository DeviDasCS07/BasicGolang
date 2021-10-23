package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main() {
	// rand.Seed(time.Now().UnixNano());
	// fmt.Println(rand.Intn(5));

	randomNum, _ := rand.Int(rand.Reader, big.NewInt(5));
	fmt.Println(randomNum)
}