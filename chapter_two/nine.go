package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"sort"
)

// How to Confirm
// cat hightemp.txt | cut -f1 | sort | uniq -c | sort
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
	fmt.Println(pref)

	a := List{}
	for k, v := range pref {
		e := Entry{k, v}
		a = append(a, e)
	}

	sort.Sort(a)
	for _, v := range a {
		fmt.Println(v)
	}
}


type Entry struct {
	name  string
	value int
}
type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].name < l[j].name)
	} else {
		return (l[i].value < l[j].value)
	}
}
