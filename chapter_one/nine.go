package main

import (
	"fmt"
	"strings"
	"math/rand"
)

func main() {
	str := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

	words := strings.Split(str, " ")
	for i, v := range words {
		str_rune := []rune(v)
		if len(str_rune) >= 4 {
			// 先頭を追加
			result := string(str_rune[0])
			// 先頭と末尾以外を抽出してランダムにシャッフル
			target := str_rune[1:(len(str_rune)-1)]
			// 末尾を追加
			result += shuffle(target) + string(str_rune[len(str_rune)-1])
			words[i] = result
		}
	}
	fmt.Println(words)
}

func shuffle(target []rune) string {
	result := ""
	for i := range target {
		j := rand.Intn(i + 1)
		target[i], target[j] = target[j], target[i]
	}
	for _, v := range target {
		result += string(v)
	}
	return result
}
