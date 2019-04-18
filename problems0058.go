package main

import "fmt"

func lengthOfLastWord(s string) int {
	size := len(s)
	result := 0
	if size == 0 {
		return 0
	}
	for end := size - 1; end >= 0; end-- {
		if s[end] == ' ' {
			if result != 0 {
				return result
			}
			continue
		}
		result++
	}
	return result

}

func main() {
	s := "HelloWord"
	fmt.Println(lengthOfLastWord(s))
}
