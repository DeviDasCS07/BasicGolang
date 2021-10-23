package main

import "fmt"

func main() {
	loginCount := 20;
	var result string
	if loginCount < 10 {
		result = "Regular user";
	} else if loginCount > 10 {
		result = "Amazing user";
	} else {
		result = "user";
	}
	fmt.Println(result);

	if num :=3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("not eq")
	}

}