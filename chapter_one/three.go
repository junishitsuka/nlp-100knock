package main

import "fmt"

func main() {
	str1 := "パトカー"
	str2 := "タクシー"
	// 一旦rune配列に変換してから文字列操作を行う
	rune1 := []rune(str1)
	rune2 := []rune(str2)
	result := ""

	for i, _ := range rune1 {
		result += string(rune1[i])
		result += string(rune2[i])
	}
	fmt.Println(result)
}
