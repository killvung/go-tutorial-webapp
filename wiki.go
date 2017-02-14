package main

import (
	"fmt"
	"io/ioutil"
    "net/http"
)

type Page struct {
    Title string 
    Body []byte //io libraries need to use byte
}

// Create persistent stroage method save() for Page
// returns a value of type error
func (p *Page) save() error{
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename,p.Body, 0600)
}

//View a wiki page by handling the URL's prefixed with /view/
func viewHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/view/"):]
    p,_ := loadPage(title)
    fmt.Fprintf(w,"<h1>%s</h1><div>%s</div>",p.Title,p.Body)
}

//Expectly, Page.save() should return nil, indicate that there is no error
//0600 means that the file should be create with read-write permission

// load the Page struct based on the title
// Of course, you need to know whether there's an error or not
func loadPage(title string) (*Page,error) {
    filename := title + ".txt"
    //Underscore "_" means that we aren't returning any value for this
    body, _ := ioutil.ReadFile(filename)
    return &Page{Title: title, Body:body},nil
}

func main(){
    http.HandleFunc("/view/",viewHandler)
    http.ListenAndServe(":8080",nil)
}
