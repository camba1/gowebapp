package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

// Page: Holds page we need to save/load
type Page struct {
	Title string
	Body  []byte
}

type fileSpec struct {
	Name     string
	ModTime  string
	Size     int64
	FullName string
}

//fileListing: List of filenames in a directory
type fileListing struct {
	Count     int
	Names     []string
	FullNames []string
	DirName   string
	Files     []fileSpec
}

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
var validFileName = regexp.MustCompile("^([a-zA-Z0-9]+)$")

// getFileName build file name base on tile + the file extension
func getFileName(title string) string {
	return filesDir + title + txtExtension
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

//loadFileListing: Get the list of files in a given directory
func loadFileListing(dirName string) ([]os.FileInfo, error) {
	fileListing, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil, err
	}
	return fileListing, nil
}

//listDirFiles: Load the count and names of the files in a directory into the fileListing struct
func listDirFiles(directory string) (fileListing, error) {
	files := fileListing{DirName: directory}
	filesInDir, err := loadFileListing(directory)
	if err != nil {
		//Do nothing
	} else {
		fileSpecs := make([]fileSpec, len(filesInDir))
		for i, info := range filesInDir {
			if !info.IsDir() {
				fileSpecs[i] = fileSpec{
					FullName: info.Name(),
					Name:     info.Name()[:len(info.Name())-4],
					ModTime:  info.ModTime().Format(dateFormat),
					Size:     info.Size(),
				}
			}
		}
		files = fileListing{
			Count:   len(fileSpecs),
			DirName: directory,
			Files:   fileSpecs,
		}
	}
	return files, err
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

//makeHandler: Take a http request to save/edit/view a file , validate urk, get file title and call the appropriate function
// Note that we use closures to pass the  function we will need to call at the end of the day as
// the first parameter to this function
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, r *http.Request) {
		match := validPath.FindStringSubmatch(r.URL.Path)
		if match == nil {
			http.NotFound(responseWriter, r)
			return
		}
		fn(responseWriter, r, match[2])
	}
}

//handler: handle http request
func handler(responseWriter http.ResponseWriter, r *http.Request) {
	_ = r
	files, err := listDirFiles(filesDir)
	err = templates.ExecuteTemplate(responseWriter, indexTemplate, files)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

//viewHandler: Handle request to view a file
func viewHandler(responseWriter http.ResponseWriter, r *http.Request, title string) {
	//title, err := getTitle(responseWriter, r)
	//if err != nil {
	//	return
	//}
	pg, err := loadPage(title)
	if err != nil {
		http.Redirect(responseWriter, r, editUrl+title, http.StatusFound)
		return
	}
	renderTemplate(responseWriter, pg, viewTemplate)

}

//editHandler: Handle request to edit a file
func editHandler(responseWriter http.ResponseWriter, r *http.Request, title string) {
	_ = r // added since we still need the handlers to be uniform so this can be called from the makeHandler function
	pg, err := loadPage(title)
	if err != nil {
		pg = &Page{Title: title}
	}
	renderTemplate(responseWriter, pg, editTemplate)
}

// saveHandler: Get data from the request body and save it into a new file
func saveHandler(responseWriter http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	pg := Page{
		Title: title,
		Body:  []byte(body),
	}
	err := pg.save()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(responseWriter, r, viewUri+title, http.StatusFound)
}

func newHandler(responseWriter http.ResponseWriter, r *http.Request) {
	match := validFileName.FindStringSubmatch(r.FormValue("fileName"))
	if match == nil {
		http.Error(responseWriter, "Invalid file name. Please use only letters and numbers and do not enter extension.", http.StatusBadRequest)
		return
	}
	title := match[1]
	http.Redirect(responseWriter, r, viewUri+title, http.StatusFound)
}

//renderTemplate: Load html template from the template cache and render page content to send back to client
func renderTemplate(responseWriter http.ResponseWriter, pg *Page, templateName string) {
	err := templates.ExecuteTemplate(responseWriter, templateName, pg)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

//startServer: Start web server and call handlers
func startServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/new/", newHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	saveDisplayTestPage()
	startServer()
}
