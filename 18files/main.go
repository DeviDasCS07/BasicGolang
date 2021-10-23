package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	writeFile();
	readFile("./test.txt");
}

func writeFile() {
	content := "Writing to a file..."
	file, err := os.Create("./test.txt");
	checkNilErr(err);
	length, err := io.WriteString(file, content);
	checkNilErr(err);
	fmt.Println(length);
	defer file.Close();
}

func readFile(fileName string){
	dataBytes, err := ioutil.ReadFile(fileName);
	checkNilErr(err);
	fmt.Println(string(dataBytes));
}

func checkNilErr(err error) {
	if err != nil {
		panic(err);
	}
}