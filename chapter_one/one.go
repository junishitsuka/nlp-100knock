package main

import "fmt"

func main() {
	target := "パタトクカシーー"

	result := ""
	for i, v := range target {
		if i % 2 == 1 {
			result += string(v)
		}
	}
	fmt.Println(result)
}

// string = []byte
// rune = unicode code point, int32 alias
// for で文字列を回すとruneが取れる
// string でキャストすれば文字列に戻る
