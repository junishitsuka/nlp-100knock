package main

import (
	"fmt"
	"os"
	"bufio"
)

// How to Confirm
// cat hightemp.txt | wc -l
func main() {
	filename := "./hightemp.txt"
	count := 0

	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		count++
	}
	fmt.Println(count)
}
