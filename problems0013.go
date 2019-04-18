package main

import "fmt"

func romanToInt(s string) int {
	// create a map that store letters to their corresponding number respectively
	result := 0
	Roman_list := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// here start the calculating method, the right number must be less than or equal to its leftside one
	// otherwise, we do the substcract method

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := Roman_list[s[i]]
		sign := 1
		if temp < last {
			sign = -1
		}
		result = result + sign*temp

		last = temp

	}
	return result
}

func main(){
	var s="IV"
	fmt.Println(romanToInt(s))
}