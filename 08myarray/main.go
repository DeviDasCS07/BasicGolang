package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "Apple";
	fruitList[1] = "Banana";
	fruitList[3] = "Peach";
	fmt.Println(fruitList);
	fmt.Println(len(fruitList));

	var vegList = [4]string{"potato","beans"};
	fmt.Println(vegList);
	fmt.Println(len(vegList));
	
}