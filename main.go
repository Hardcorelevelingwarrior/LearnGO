package main

import (
	"fmt"
	
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/main/views"
)
var (
	homeView *views.View
 	contactView *views.View
 	faqView *views.View
	signupView *views.View
)


func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	must(homeView.Render(w,nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	must(contactView.Render(w,nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	must(faqView.Render(w,nil))
}
func signup(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","text/html")
	must(signupView.Render(w,nil))
}

func notfound404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h1>404 This page not exist</h1>")
}

func main() {
	homeView = views.NewView("bootstrap","views/home.gohtml")
	contactView = views.NewView("bootstrap","views/contact.gohtml")
	faqView = views.NewView("bootstrap","views/faq.gohtml")
	signupView = views.NewView("bootstrap","views/signup.gohtml")
	var h http.Handler = http.HandlerFunc(notfound404)
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/signup",signup)
	r.NotFoundHandler = h
	http.ListenAndServe(":3000", r)
}

// A helper function that panics on any error
func must(err error) {
	if err != nil {
	panic(err)
	}
	}
//Using other router package
/*func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
}*/
//6.6.1 Ex1 - Create an FAQ page with the Bootstrap layout
//6.6.2 Ex2 - Update the navbar to link to the FAQ page