package main

import ( 
"fmt"
"net/http"
)

/* This simple webserver based on tutorial at https://youtu.be/G3PvTWRIhZA
*/
func index_handler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"Simple Example") }

func hello_handler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"Hello my name is Elliot. \n This is my website.") }

func main() { 
	http.HandleFunc("/", index_handler) 
	http.HandleFunc("/hello",hello_handler) 
	http.ListenAndServe(":8000", nil) 
}
