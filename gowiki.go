package main

import (
	"os"
	"log"
	"fmt"
)

type Page struct {
	Title string

	// We declare the Body element of type []byte because we'll use os.WriteFile() which takes 
	// []byte as its argument
	Body []byte
}

// The save() method returns an error value because that is the return type of os.WriteFile().
// If nothing wrong, the save() method will return nil
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{ Title: title, Body: []byte(body) }, nil
}

func main() {
	// Uncomment the following lines for testing save() method
	// p := &Page{ Title: "test", Body: []byte("test")}
	// p.save()

	// Uncomment the following lines for testing loadPage() method
	log.SetPrefix("gowiki: ")
	p, err := loadPage("test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Title: %s, Body: %s", p.Title, p.Body)
	
}