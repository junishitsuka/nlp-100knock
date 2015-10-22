package main

import (
	"fmt"
	"os"
	"bufio"
)

// How to Confirm
// paste col1.txt col2.txt
func main() {
	fp1, err := os.Open("col1.txt")
	if err != nil {
		panic(err)
	}
	defer fp1.Close()

	fp2, err := os.Open("col2.txt")
	if err != nil {
		panic(err)
	}
	defer fp2.Close()

	scanner1 := bufio.NewScanner(fp1)
	scanner2 := bufio.NewScanner(fp2)
	for scanner1.Scan() {
		scanner2.Scan()
		fmt.Println(scanner1.Text() + "\t" + scanner2.Text())
	}
}
