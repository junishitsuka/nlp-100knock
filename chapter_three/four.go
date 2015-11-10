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
		re := regexp.MustCompile(`\[\[File:(.+)\]\]`)
		section := re.FindAllStringSubmatch(v, -1)
		if len(section) != 0 {
			fmt.Println(section[0][0])
		}
		// for _, sec := range section {
		// 	fmt.Println(sec)
		// }
	}
}
