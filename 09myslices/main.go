package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices");
	var fruitList = []string{"Apple", "Tomato", "Peach"};
	fmt.Printf("Type of variable is %T \n", fruitList);
	
	fruitList = append(fruitList,"Mango", "Banana");

	fmt.Println(fruitList);

	fruitList = append(fruitList[:3]);
	fmt.Println(fruitList);

	highScores  := make([] int, 4);
	highScores[0] = 200;
	highScores[1] = 800;
	highScores[2] = 400;
	highScores[3] = 600;
  //highScores[4] = 900;
	highScores = append(highScores, 900, 1000, 2000);
	fmt.Println(highScores);

	sort.Ints(highScores);
	fmt.Println(highScores);
	fmt.Println(sort.IntsAreSorted(highScores));

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "js", "swift", "python", "ruby"};
	fmt.Println(courses);
	var index int = 2;
	courses = append(courses[:index], courses[index+1:]...);
	fmt.Println(courses);
}