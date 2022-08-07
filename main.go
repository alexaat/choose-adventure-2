package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var book Book
var templ *template.Template

func main() {

	templ = template.Must(template.New("").Parse(defaultPageTemplate))

	book = getBook()
	if book == nil {
		return
	}
	log.Fatal(http.ListenAndServe(":8080", requestHandler()))
}

func requestHandler() http.Handler {
	handler := Handler{book: book}
	return handler
}

type Handler struct {
	book Book
}

func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var path = strings.TrimPrefix(r.URL.Path, "/")
	if path == "" || path == "/" {
		page := book["intro"]
		templ.Execute(w, page)
	}
	if page, ok := book[path]; ok {
		templ.Execute(w, page)
	}
}

func getBook() Book {
	data, err := readFile("gopher.json")
	if err != nil {
		fmt.Println("Error reading file")
		return nil
	}

	book, err := decodeJson(data)
	if err != nil {
		fmt.Println("Error decoding json")
		return nil
	}
	return book
}
