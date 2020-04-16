package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// Page: Holds page we need to save/load
type Page struct {
	Title string
	Body  []byte
}

// File manipulation constants
const (
	txtExtension      string = ".txt"
	readWriteFileMode        = 0600
)

// Html templates
const (
	editTemplate = "edit.html"
	viewTemplate = "view.html"
)

const (
	editUrl = "/edit/"
	saveUri = "/save/"
	viewUri = "/view/"
)

//templates: Html templates cache loaded from disk
var templates = template.Must(template.ParseFiles(editTemplate, viewTemplate))

//validPath: Regular expression to be validate that we have a valid path to save and retrieve files
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// getFileName build file name base on tile + the file extension
func getFileName(title string) string {
	return title + txtExtension
}

//save: Save a page to disk
func (pg *Page) save() error {
	fileName := getFileName(pg.Title)
	return ioutil.WriteFile(fileName, pg.Body, readWriteFileMode)

}

//loadPage: load page from disk into a page struct
func loadPage(title string) (*Page, error) {
	fileName := getFileName(title)
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{title, body}, nil
}

//saveDisplayTestPage: Save a string into a file. Reload the file and display content.
func saveDisplayTestPage() {
	pageTitle := "History"
	pg := Page{
		Title: pageTitle,
		Body:  []byte("This is history!"),
	}
	err := pg.save()
	if err != nil {
		panic(err)
	}
	pgLd, err := loadPage(pageTitle)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(pgLd.Title + ": " + string(pgLd.Body))
	}
}

func getTitle(responseWriter http.ResponseWriter, r *http.Request) (string, error) {
	match := validPath.FindStringSubmatch(r.URL.Path)
	if match == nil {
		http.NotFound(responseWriter, r)
		return "", errors.New("invalid page title")
	}
	return match[2], nil
}

//handler: handle http request
func handler(responseWriter http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(responseWriter, "helle from %s", r.URL.Path)
	if err != nil {
		return
	}
}

//viewHandler: Handle request to view a file
func viewHandler(responseWriter http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len(viewUri):]
	title, err := getTitle(responseWriter, r)
	if err != nil {
		return
	}
	pg, err := loadPage(title)
	if err != nil {
		http.Redirect(responseWriter, r, editUrl+title, http.StatusFound)
		return
	}
	renderTemplate(responseWriter, pg, viewTemplate)

}

//editHandler: Handle request to edit a file
func editHandler(responseWriter http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len(editUrl):]
	title, err := getTitle(responseWriter, r)
	if err != nil {
		return
	}
	pg, err := loadPage(title)
	if err != nil {
		pg = &Page{Title: title}
	}
	renderTemplate(responseWriter, pg, editTemplate)
}

func saveHandler(responseWriter http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len(saveUri):]
	title, err := getTitle(responseWriter, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	pg := Page{
		Title: title,
		Body:  []byte(body),
	}
	err = pg.save()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(responseWriter, r, viewUri+title, http.StatusFound)
}

//renderTemplate: Load html template from the template cache and render page content to send back to client
func renderTemplate(responseWriter http.ResponseWriter, pg *Page, templateName string) {
	//t, err := template.ParseFiles(templateName)
	//if err != nil {
	//	http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//err = t.Execute(responseWriter, pg)
	err := templates.ExecuteTemplate(responseWriter, templateName, pg)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

//startServer: Start web server and call handlers
func startServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	saveDisplayTestPage()
	startServer()
}
