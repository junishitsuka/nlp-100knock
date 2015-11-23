package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"os"
	"net/http"
)

// リクエストの結果を文字列で返す関数
func execRequest(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)
}

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
		re := regexp.MustCompile(`\[\[ファイル:(.+)\]\]`)
		section := re.FindAllStringSubmatch(v, -1)
		if len(section) != 0 {
			// ファイル名部分を取得
			file := strings.Split(section[0][1], "|")
			// 空白スペースをエンコード
			uri := strings.Replace(file[0], " ", "%20", -1)

			api := "https://commons.wikimedia.org/w/api.php?action=query&format=json&titles=File:%s&prop=imageinfo&iiprop=url"
			url := fmt.Sprintf(api, uri)
			fmt.Println(url)

			fmt.Println(execRequest(url))
			os.Exit(0)
		}
	}
}
