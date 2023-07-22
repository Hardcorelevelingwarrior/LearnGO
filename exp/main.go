package main

import (
	"html/template"
	"os"
)

// Ex1: Add new variable to hello.gohtml
// 1. Adding a new field to the data variable in your go code
// 2. Using that new field inside of your template
func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	data := struct {
		Name     string
		Age      int
		Weight   float64
		Favorite map[string]bool
	}{"Minh", 20, 60.05, map[string]bool{"fat": true}}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

//Ex2 : Experiment with different data types
//Ex3 - [HARD] Learn how to use nested data : access by key : .Favorite.food
//Ex4 - [HARD] Create an if/else statement in your tem-plate

//5.3.1 Ex1 - What does MVC stand for? 
//M stand for model, v stand for views and c for controller
//5.3.2 Ex2 - What is each layer of MVC responsible for?
//Views to render the output, controller for determine what model to use or what template to use , model is the part where we code our logic
//5.3.3 Ex3 - What are some benefits to using MVC?
//To structure our code for better read and maintain