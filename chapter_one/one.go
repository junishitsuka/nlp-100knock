package main

import "fmt"

func main() {
	target := "stressed"

	result := ""
	l := len(target) // len(str) で文字列・配列の長さを取得できる
	for i, _ := range target {
		result += string(target[l - i - 1]) // + 演算子で文字列を結合
	}
	fmt.Println(result)
}
