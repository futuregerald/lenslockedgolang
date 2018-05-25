package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Message", "Gerald made me!")
	fmt.Fprint(w, "<h2>Is Gerald Awesome?</h2> <br /> <p>Why yes, yes he is</p>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Message", "Gerald is still awesome")
	fmt.Fprint(w, "<h2>Gerald can't find what you're looking for. Are you sure you have the right URL?</p>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")
}

func main() {
	fmt.Println("The server is running!")
	var h http.Handler = http.HandlerFunc(notFound)
	r := mux.NewRouter()
	r.NotFoundHandler = h
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}
