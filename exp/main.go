package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	type info struct {
		Town    string
		ZipCode int
	}

	Random := []string{"fries", "turkey", "drinks"}
	/*
		data := struct {
			Name string
			Age  int
			Info info
		}{"Gerald", 32, info{"New York", 10004}}
	*/
	err = t.Execute(os.Stdout, Random)
	if err != nil {
		panic(err)
	}
}
