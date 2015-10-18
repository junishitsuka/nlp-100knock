package main

import (
	"fmt"
	"strings"
)

// 関数の型は基本省略できない
// 引数が同じ型なら最後にまとめて宣言できる
func arrayExists(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	result := make(map[string]int)
	statement := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	statement = strings.Trim(statement, ".")
	target := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
	words := strings.Split(statement, " ")

	for i, v := range words {
		runes := []rune(v)
		if arrayExists(target, i) {
			key := string(runes[:2])
			result[key] = i
		} else {
			key := string(runes[:1])
			result[key] = i
		}
	}

	fmt.Println(result)
}
