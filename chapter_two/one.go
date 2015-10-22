package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

// How to Confirm
// cat hightemp.txt | sed -e "s/\t/ /g"
func main() {
	filename := "./hightemp.txt"

	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		replaced := strings.Replace(scanner.Text(), "\t", " ", -1)
		fmt.Println(replaced)
	}
}
