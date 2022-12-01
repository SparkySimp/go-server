package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

func (self *Page) save() error {
	filename := self.Title + ".html"
	return os.WriteFile(filename, self.Body, 0600)
}

func loadPage(title string) (*Page, error) {
		filename := title + ".txt"
		body, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		return &Page{Title: title, Body: body}, nil
}

func (self *Page) Render() string {
	inlineHtml := "<!DOCTYPE html> \n\r"
	inlineHtml += "<html>\n\r"
	inlineHtml += "\t<head>\n\r"
	inlineHtml += "\t\t<title>\n\r"
	inlineHtml += "\t\t\t" + self.Title + "\n\r"
	inlineHtml += "\t\t</title>\n\r"
	inlineHtml += "\t</head>\n\r"
	inlineHtml += "\t<body>\n\r"
	inlineHtml += "%v"
	inlineHtml += "\t</body>\n\r"
	inlineHtml += "</html>"
	return fmt.Sprintf(inlineHtml, string(self.Body))
}

func main() {
	page := Page{ "Hello World!", []byte("<p>Example page</p>") }
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, page.Render())
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}