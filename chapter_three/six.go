package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
)

func main() {
	filename := "./jawiki-british.json"
	result := make(map[string]string)

	// ファイル全体を読み込む
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	text := string(buf)

	// 基礎情報にマッチする部分を抽出
	re := regexp.MustCompile(`\{\{基礎情報 (.+)\}\}`)
	section := re.FindAllStringSubmatch(text, -1)

	// 基礎情報の区切りで配列に変換する
	infoDic := strings.Split(section[0][0], "\\n|")

	// key と value を取得して辞書に入れる
	for _, v := range infoDic {
		re := regexp.MustCompile(`(.+) = (.+)`)
		info := re.FindAllStringSubmatch(v, -1)
		if len(info) >  0 {
			result[string(info[0][1])] = string(removeEmphasis(info[0][2]))
		}
	}

	for k, v := range result {
		fmt.Println(k + " = " + v)
	}
}

func removeEmphasis(target string) string {
	tmp1 := strings.Replace(target, "'''''", "", -1)
	tmp2 := strings.Replace(tmp1, "'''", "", -1)
	removed := strings.Replace(tmp2, "''", "", -1)
	return removed
}
