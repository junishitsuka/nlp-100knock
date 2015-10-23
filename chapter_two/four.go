package main

import (
	"fmt"
	"os"
	"bufio"
	"flag"
)

// How to Confirm
// cat hightemp.txt | head -n N
func main() {
	fp, err := os.Open("hightemp.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// get args
	n := flag.Int("n", 0, "input natural number")
	flag.Parse()

	scanner := bufio.NewScanner(fp)
	for i := 0; i < *n; i++ {
		scanner.Scan()
		fmt.Println(scanner.Text())
	}
}
