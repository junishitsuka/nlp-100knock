package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

// How to Confirm
// cat hightemp.txt | cut -f1
// cat hightemp.txt | cut -f2
func main() {
	filename := "./hightemp.txt"

	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	fp1, err := os.Create("col1.txt")
	if err != nil {
		panic(err)
	}
	defer fp1.Close()
	fw1 := bufio.NewWriter(fp1)

	fp2, err := os.Create("col2.txt")
	if err != nil {
		panic(err)
	}
	defer fp2.Close()
	fw2 := bufio.NewWriter(fp2)

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), "\t")
		fmt.Fprint(fw1, splitted[0] + "\n")
		fmt.Fprint(fw2, splitted[1] + "\n")
		fw1.Flush()
		fw2.Flush()
	}
}
