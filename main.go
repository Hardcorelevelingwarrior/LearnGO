package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*
	func home(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html")
		fmt.Fprint(w, "<h1>Welcome</h1>")
	}

	func contact(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html")
		fmt.Fprint(w, "To get in touch, please send an email "+
			"to <a href=\"mailto:support@lenslocked.com\">"+
			"support@lenslocked.com</a>.")
	}

	func faq(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html")
		fmt.Fprint(w, "<h1>This is our faq page</h1>")
	}

	func notfound404(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html")
		fmt.Fprint(w, "<h1>404 This page not exist</h1>")
	}

	func main() {
		var h http.Handler = http.HandlerFunc(notfound404)
		r := mux.NewRouter()
		r.HandleFunc("/", home)
		r.HandleFunc("/contact", contact)
		r.HandleFunc("/faq", faq)
		r.NotFoundHandler = h
		http.ListenAndServe(":3000", r)
	}
*/
func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h1>Welcome</h1>")
}
func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")
}
func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h1>This is our faq page</h1>")
}
func notfound404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h1>404 This page not exist</h1>")
}
func main() {
	var h http.Handler = http.HandlerFunc(notfound404)
	router := httprouter.New()
	router.GET("/", home)
	router.GET("/contact", contact)
	router.GET("/faq", faq)
	router.NotFound = h
	http.ListenAndServe(":3000", router)
}
