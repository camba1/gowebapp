package main

import (
	"./fileManager"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

//dateFormat: Formatting to display dates in the screen
const dateFormat string = "Jan-02-06"

// File manipulation constants
const (
	txtExtension      string = ".txt"
	readWriteFileMode        = 0600
	filesDir                 = "files/"
)

// Html templates
const (
	templatesDir  = "htmlTemplates/"
	editTemplate  = "edit.html"
	viewTemplate  = "view.html"
	indexTemplate = "index.html"
)

// Page uris
const (
	editUrl = "/edit/"
	//saveUri = "/save/"
	viewUri = "/view/"
)

//templates: Html templates cache loaded from disk
var templates = template.Must(template.ParseFiles(templatesDir+editTemplate, templatesDir+viewTemplate, templatesDir+indexTemplate))

//validPath: Regular expression to be validate that we have a valid path to save and retrieve files
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

//validFileName: Validate that the file name only contains letters and numbers
var validFileName = regexp.MustCompile("^([a-zA-Z0-9]+)$")

//startServer: Start web server and call handlers
func startServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/new/", newHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/healthCheck", healthCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fileManager.SaveDisplayTestPage(filesDir, txtExtension, readWriteFileMode)
	startServer()
}
