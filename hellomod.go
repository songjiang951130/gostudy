package main

import (
	"fmt"
	"net/http"
)

func Indexhandler(w http.ResponseWriter,r *http.Request)  {
    fmt.Fprintln(w,"hello world")
}

func main() {
	fmt.Println("hello,time99")
	var val int = 10
	fmt.Println(val)
	http.HandleFunc("/",Indexhandler)
    http.ListenAndServe("127.0.0.1",nil)
}
