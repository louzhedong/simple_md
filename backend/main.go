package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("hello")
}

func main()  {
	http.HandleFunc("/", sayHello)
	fmt.Println("Listen port 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}