package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"net/http"
)

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")  // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true") //设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") //允许请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8")             //返回数据格式是json
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		f(w, r)
	}
}

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
	http.HandleFunc("/md", cors(mdDemo))
	fmt.Println("Listen port 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}