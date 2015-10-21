package main

import "fmt"

func main() {
	str := "password"
	result := []rune{}
	for _, v := range str {
		if v >= 96 && v < 123 { 
			result = append(result, 219 - v)
		}
	}
	fmt.Println(string(result))
}
