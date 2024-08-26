package main

import "net/http"

func main() {
	router := GetRouter()

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
