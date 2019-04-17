package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	len1, len2 := len(haystack), len(needle)
	for i := 0; i <= len1-len2; i++ {
		if haystack[i:i+len2] == needle {
			return i
		}
	}
	return -1

}

func main() {
	var result1, result2 int
	result1 = strStr("hello", "ll")
	result2 = strStr("aaaaa", "bba")
	fmt.Println(result1, result2)
}
