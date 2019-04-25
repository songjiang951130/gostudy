package main

import (
	"fmt"
	"net/http"
)

func Indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world 222")
}

func main() {
	fmt.Println("hello,time99")
	var val int = 10
	fmt.Println(val)
	digits := "23"
	// var result [] string =
	letterCombinations(digits)
	// for i := 0; i < len(result); i++ {
	//     fmt.Println(result[i])
	// }

}

func letterCombinations(digits string) {
	fmt.Println(len(digits))
	distance := 47
	result := make([]string, len(digits)*3)
	fmt.Println(result)
	for j := 0; j < len(digits); j++ {
		var temp string
		for index := 0; index < 3; index++ {
			temp = temp + string(digits[j]+byte(distance+index))
			fmt.Println("j:", j, string(digits[j]+byte(distance+index)))
		}
		fmt.Println("temp:", temp)

	}
}
