package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano());
	diceNumber := rand.Intn(6)+1;
	fmt.Println(diceNumber);

	switch diceNumber {
	case 1:
		fmt.Println("Dice is open and value is 1");
	case 2:
		fmt.Println("Dice value is 2 and move 2 spot");
	case 3:
		fmt.Println("Dice value is 3 and move 3 spot");
		fallthrough;
	case 4:
		fmt.Println("Dice value is 4 and move 4 spot");
		fallthrough;
	case 5:
		fmt.Println("Dice value is 5 and move 5 spot");
	case 6:
		fmt.Println("Dice value is 6 and roll again");			
	default:
		fmt.Println("...");	
	}
}