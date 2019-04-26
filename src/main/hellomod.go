package main

import (
	// "container/list"
	"fmt"
	"strconv"
)

func main() {
	digits := "24"
	result := letterCombinations(digits)
	fmt.Println(result)
}

func letterCombinations(digits string) []string {
	numberMapping := make(map[int]string, 8)
	numberMapping[2] = "abc"
	numberMapping[3] = "def"
	numberMapping[4] = "ghi"
	numberMapping[5] = "jkl"
	numberMapping[6] = "mno"
	numberMapping[7] = "pqrs"
	numberMapping[8] = "tuv"
	numberMapping[9] = "wxyz"
	fmt.Println(len(digits))
	result := make([]string, 0)
	fmt.Println(len(result), cap(result))

	dfs(result, digits, numberMapping, "")
	return result
}

func dfs(result []string, digits string, numberMapping map[int]string, tmp string) {
	if len(tmp) == len(digits) {
		result = append(result, tmp)
	}
	for _, ch := range digits {
		key, _ := strconv.Atoi(string(ch))
		value := numberMapping[key]
		fmt.Println(key, value)
	}

}
