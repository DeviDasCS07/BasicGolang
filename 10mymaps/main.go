package main

import "fmt"

func main() {
	languages := make(map[string]string);
	languages["JS"] = "javascript";
	languages["RB"] = "ruby";
	languages["Swift"] = "swift";
	fmt.Println(languages);
	fmt.Println(languages["JS"]);
	delete(languages, "RB");
	fmt.Println(languages);

	// iterate through maps
	for key, value := range languages {
		fmt.Printf("Key %v -> Value %v\n", key, value);
	}
}