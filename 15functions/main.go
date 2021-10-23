package main

import "fmt"

func main() {
	msg, result := adder(1,2,3,4,5,6,7,8,9,10);
	fmt.Println(msg);
	fmt.Println(result);
}

func adder(values...int) (string, int) {
	total := 0;
	for _,val := range values {
		total += val;
	}
	return "Total",total;
}
