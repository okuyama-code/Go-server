package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
fmt.Fprintf(w, "hello")
}


func main() {
	// "./static" ディレクトリからファイルを提供するファイルサーバを作成
	fileServer := http.FileServer(http.Dir("./static"))

	// ルートURL ("/") へのリクエストをファイルサーバに処理させる
	http.Handle("/", fileServer)

	// "/form" へのリクエストを formHandler 関数で処理する
	http.HandleFunc("/form", formHandler)

	// "/hello" へのリクエストを helloHandler 関数で処理する
	http.HandleFunc("/hello", helloHandler)

	// サーバを指定したポートで起動する
	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}




