package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"sort"
)

type stringSlice []string

func (p stringSlice) Len() int           { return len(p) }
func (p stringSlice) Less(i, j int) bool {
	split_i := strings.Split(p[i], "\t")
	split_j := strings.Split(p[j], "\t")
	return split_i[2] < split_j[2]
}
func (p stringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// How to Confirm
// cat hightemp.txt | sort -k 3 
func main() {
	filename := "./hightemp.txt"

	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	pref := []string{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		pref = append(pref, scanner.Text())
	}
	sort.Sort(stringSlice(pref))
	for _, v := range pref {
		fmt.Println(v)
	}
}
