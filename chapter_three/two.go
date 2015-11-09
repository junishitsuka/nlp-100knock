package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
)

func main() {
	filename := "./jawiki-british.json"

	// ファイル全体を読み込む
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := string(buf)
	line := strings.Split(text, "\\n")

	for _, v := range line {
		if m, _ := regexp.MatchString("Category", v); m {
			fmt.Println(v)
		}
	}
}
