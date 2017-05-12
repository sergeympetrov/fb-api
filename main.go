package main

import (
	"fmt"
	"time"
	"net/http"
	"html/template"
)

type page struct {
	Title string	//имена переменных должны совпадать с тем, что мы написали выше!
	Msg string		//и переменные обязательно должны быть публичными!
	Path string
}

func testPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")

	titleR := r.URL.Path[len("/"):]
	fmt.Println(titleR, time.Now())

	var Message string = "Hello World, time"+"sdf"


	t, _ := template.ParseFiles("index.html")
	t.Execute(w, &page{Title: "Just page", Msg: Message, Path: titleR })
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	//w.Write([]byte("hi again..."))
	fmt.Println("reload!")
	fmt.Fprint(w, "<html><body><h1>hi, there!</h1><p>"+r.URL.Path+"</p></body></html>")
}

func main() {
	//fmt.Println(*http.Request.URL.Path)

	http.HandleFunc("/sdf", index)
	http.HandleFunc("/test/", testPage)

	http.ListenAndServe(":8011", nil)
}