package main

import (
	"fmt"
)

// 文字列の配列から要素のNgramを作成する関数
func makeNgram(arr []string, n int) (result []string) {
	// validation
	if len(arr) < n {
		return arr
	}

	for index := 0; index + n - 1 < len(arr) ; index++ {
		ngram := ""
		for i := index; i < n + index; i++ {
			ngram += arr[i]
		}
		result = append(result, ngram)
	}
	return
}

func arrayExists(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// 戻り値に名前を付けておくとreturnに変数名を書かなくていい
// 宣言も不要
func makeSet(arr []string) (result []string) {
	for _, v := range arr {
		if !arrayExists(result, v) {
			result = append(result, v)
		}
	}
	return
}

func sumSet(arr1 []string, arr2 []string) []string {
	result := arr1

	for _, v := range arr2 {
		if !arrayExists(arr1, v) {
			result = append(result, v)
		}
	}
	return result
}

func intersectionSet(arr1 []string, arr2 []string) (result []string) {
	for _, v1 := range arr1 {
		if arrayExists(arr2, v1) {
			result = append(result, v1)
		}
	}
	return
}

// arr2 - arr1
func differenceSet(arr1 []string, arr2 []string) (result []string) {
	for _, v2 := range arr2 {
		if !arrayExists(arr1, v2) {
			result = append(result, v2)
		}
	}
	return
}

func checkStr(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	str1 := "paragraph"
	str2 := "raparaparadise"

	arr1 := []string{}
	arr2 := []string{}

	for _, v := range str1 {
		arr1 = append(arr1, string(v))
	}
	for _, v := range str2 {
		arr2 = append(arr2, string(v))
	}

	// 文字bi-gram集合
	bigram1 := makeSet(makeNgram(arr1, 2))
	bigram2 := makeSet(makeNgram(arr2, 2))
	fmt.Println(bigram1, bigram2)

	// 和集合
	fmt.Println(sumSet(bigram1, bigram2))
	// 積集合
	fmt.Println(intersectionSet(bigram1, bigram2))
	// 差集合
	fmt.Println(differenceSet(bigram1, bigram2))

	// "se"が含まれるかどうか
	if checkStr(bigram1, "se") {
		fmt.Println("X: true")
	} else {
		fmt.Println("X: false")
	}
	if checkStr(bigram2, "se") {
		fmt.Println("Y: true")
	} else {
		fmt.Println("X: false")
	}
}
