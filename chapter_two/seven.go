package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

// How to Confirm
// cat hightemp.txt | cut -f1 | sort | uniq | wc -l
func main() {
	filename := "./hightemp.txt"

	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	pref := make(map[string]int)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), "\t")
		if _, ok := pref[splitted[0]]; ok {
			pref[splitted[0]] += 1
		} else {
			pref[splitted[0]] = 1
		}
	}
	fmt.Println(len(pref))
}
