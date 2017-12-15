package main

import (
	"net/http"
	"fmt"
)

func main() {

	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)

	http.HandleFunc("/layout", handleTemplate)

	fmt.Println("Listening on port 8080.....")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}

