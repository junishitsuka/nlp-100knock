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

	sentences := []string{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		sentences = append(sentences, scanner.Text())
	}
	for i := len(sentences) - *n; i < len(sentences); i++ {
		fmt.Println(sentences[i])
	}
}
