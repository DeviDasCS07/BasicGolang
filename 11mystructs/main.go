package main

import "fmt"

func main() {
	// no inheritance, no super or parent
	user := User{"AAA", "AAA@gmail.com", true, 25}
	fmt.Println(user)
	fmt.Printf("Details %+v\n", user);
	fmt.Printf("Name %v and Email %v", user.Name, user.Email);
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}