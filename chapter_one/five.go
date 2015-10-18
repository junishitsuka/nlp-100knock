package main

import (
	"fmt"
	"strings"
)

// 文字列の配列から要素のNgramを作成する関数
func makeNgram(arr []string, n int) []string {
	// validation
	if len(arr) < n {
		return arr
	}

	result := []string{}

	for index := 0; index + n - 1 < len(arr) ; index++ {
		ngram := ""
		for i := index; i < n + index; i++ {
			ngram += arr[i]
		}
		result = append(result, ngram)
	}
	return result
}

func main() {
	target := "I am an NLPer"
	target_array := strings.Split(target, " ")

	// 単語bi-gram
	fmt.Println(makeNgram(target_array, 2))

	// 文字bi-gram
	string_array := []string{}

	// Trimは先頭か末尾しか取り除かないので, strings.Replaceを使用する
	for _, v := range strings.Replace(target, " ", "", -1) {
		string_array = append(string_array, string(v))
	}
	fmt.Println(makeNgram(string_array, 2))
}
