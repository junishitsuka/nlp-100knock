package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/json"
)

type Wiki struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func main() {
	filename := "./jawiki-country.json"

	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		var wiki Wiki
		text := scanner.Text()
		line := []byte(text)
		err := json.Unmarshal(line, &wiki)
		if err != nil {
			panic(err)
		}

		fmt.Println(wiki.Title)
		if wiki.Title == "イギリス" {
			fmt.Println(wiki.Text)
		}
	}
}
