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
