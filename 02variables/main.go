package main

import "fmt"

func main() {
	var user string = "testuser";
	fmt.Println(user);
	fmt.Printf("type of variable : %T ", user);
	fmt.Println("\n-----------");
	
	var isLoggedIn bool = false;
	fmt.Println(isLoggedIn);
	fmt.Printf("type of variable : %T ", isLoggedIn);
	fmt.Println("\n-----------");

	var smallVal int = 256;
	fmt.Println(smallVal);
	fmt.Printf("type of variable : %T ", smallVal);
	fmt.Println("\n-----------");

	var floatVal float64 = 256.111321321321313132;
	fmt.Println(floatVal);
	fmt.Printf("type of variable : %T ", floatVal);
	fmt.Println("\n-----------");

	var tempVal int;
	fmt.Println(tempVal);
	fmt.Printf("type of variable : %T ", tempVal);
	fmt.Println("\n-----------");

	
	var webSite = "www.google.com";
	fmt.Println(webSite);

	fmt.Println("\n-----------");

	numberOfUsers := 10;
	fmt.Println(numberOfUsers);
}