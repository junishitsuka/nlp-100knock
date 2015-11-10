package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
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
		re := regexp.MustCompile(`==(.{0,20})==`)
		section := re.FindAllStringSubmatch(v, -1)
		for _, sec := range section {
			level := strings.Count(sec[0], "=") / 2 - 1
			fmt.Println(strings.Replace(sec[0], "=", "", -1) + ": " + strconv.Itoa(level))
		}
	}
}
