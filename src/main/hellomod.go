package main

import (
	"fmt"
)

func main() {
	digits := "24"

	letterCombinations(digits)
}

func letterCombinations(digits string) {
	numberMapping := make([]string, 10)
	numberMapping[2] = "abc"
	numberMapping[3] = "def"
	numberMapping[4] = "ghi"
	numberMapping[5] = "jkl"
	numberMapping[6] = "mno"
	numberMapping[7] = "pqrs"
	numberMapping[8] = "tuv"
	numberMapping[9] = "wxyz"
	fmt.Println(len(digits))
	distance := 47
	result := make([]string, len(digits)*3)
	fmt.Println(result)
	for j := 0; j < len(digits); j++ {
		// for index := 0; index < 3; index++ {
		var temp string
		for l := 0; l < len(digits); l++ {
			temp = temp + string(digits[l]+byte(distance))
		}
		fmt.Println("j:", j, string(digits[j]+byte(distance)))
		fmt.Println("temp:", temp)
		// }
	}
}

func num2char(number int) []byte {
	result := make([]byte, 3)
	distance := 47
	for index := 0; index < len(result); index++ {
		result[index] = byte(number + distance)
	}
	return result
}
