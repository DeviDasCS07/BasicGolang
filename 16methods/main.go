package main

import "fmt"

func main() {
	// no inheritance, no super or parent
	user := User{"AAA", "AAA@gmail.com", true, 25}
	fmt.Println(user)
	fmt.Printf("Details %+v\n", user);
	fmt.Printf("Name %v and Email %v\n", user.Name, user.Email);
	user.GetStatus();
	user.CreateMail();
	fmt.Printf("Name %v and Email %v\n", user.Name, user.Email);
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println(u.Status);
}

func (u User) CreateMail() {
	u.Email = "test@gmail.com";
	fmt.Println(u.Email);
}