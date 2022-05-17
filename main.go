package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	dotenv "github.com/joho/godotenv"
)

type Page struct {
	Title string
	Body []byte
}

// TODO What is the *Struct and &Struct doing?
func (p *Page) save() error {
	filename := p.Title + ".txt"
	// fmt.Println(filename)
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Println(title)
	// p, _ := loadPage(title)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	dotenv.Load()
	dbName := os.Getenv("DATABASE_NAME")
	fmt.Println(dbName)
	// p1 := &Page {Title: "TestPage", Body: []byte("this is a sample.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
	// http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
