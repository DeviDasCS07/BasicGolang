package main

import "fmt"

func main() {
	// var ptr *int;
	// fmt.Println(ptr)
	number := 20;
	var ptr = &number;
	fmt.Println(ptr);
	fmt.Println(*ptr);	
	*ptr = *ptr + 5;
	fmt.Println(number);
}