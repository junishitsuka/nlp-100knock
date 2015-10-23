package main

import (
	"fmt"
	"os"
	"bufio"
	"flag"
	"strconv"
)

// How to Confirm
// split hightemp.txt -l N
func main() {
	fp, err := os.Open("hightemp.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// get args
	n := flag.Int("n", 1, "input natural number")
	flag.Parse()

	scanner := bufio.NewScanner(fp)
	index := 0
	for scanner.Scan() {
		fp, err := os.Create("split_" + strconv.Itoa(index) + ".txt")
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fw := bufio.NewWriter(fp)

		for i := 0; i < *n; i++ {
			fmt.Fprint(fw, scanner.Text() + "\n")
			if (i != *n - 1) {
				scanner.Scan()
			}
		}
		index++
		fw.Flush()
	}
}
