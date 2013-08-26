package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var Config = map[string]string{
	"root": "/home/deepak/Public",
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := Config["root"] + p.Title + ".html"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func load(title string) (*Page, error) {
	filename := Config["root"] + "/" + title
	stat, _ := os.Stat(filename)
	if stat != nil && stat.IsDir() {
		filename = filename + "/index.html"
	} else {
		filename = filename + ".html"
	}

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func handleRoute(route string, w http.ResponseWriter, method string) {
	if route == "" {
		route = "index"
	}

	if strings.ToLower(method) != "get" {
		w.WriteHeader(http.StatusBadRequest)
		route = "400"
	}

	p, _ := load(route)

	if p != nil {
		fmt.Fprintf(w, "%s", string(p.Body))
	} else {
		p, _ = load("404")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, string(p.Body))
	}

}

func router(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.Proto, r.URL)
	route := r.URL.Path[1:]
	handleRoute(route, w, r.Method)
}

func main() {
	http.HandleFunc("/", router)
	http.ListenAndServe(":8000", nil)
}
