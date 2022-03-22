package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("hello")
}

func mdDemo(w http.ResponseWriter, r *http.Request) {
	filepath := "./resource/markdown/vscode_plugin.md"
	input, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	unsafe := blackfriday.MarkdownCommon(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	w.Write(html)
	fmt.Println(html)
}

func main()  {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/md", mdDemo)
	fmt.Println("Listen port 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}