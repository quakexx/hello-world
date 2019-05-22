package main

import "net/http"

func main(){
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello2", hello2)
	http.ListenAndServe(":8080",nil)
}

func hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello!"))
}

func hello2(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello2!"))
}
