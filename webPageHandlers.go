package main

import (
	"./fileManager"
	"io"
	"net/http"
)

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
	files, err := fileManager.ListDirFiles(filesDir, dateFormat)
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
	pg, err := fileManager.LoadPage(title, filesDir, txtExtension)
	if err != nil {
		http.Redirect(responseWriter, r, editUrl+title, http.StatusFound)
		return
	}
	renderTemplate(responseWriter, pg, viewTemplate)

}

//editHandler: Handle request to edit a file
func editHandler(responseWriter http.ResponseWriter, r *http.Request, title string) {
	_ = r // added since we still need the handlers to be uniform so this can be called from the makeHandler function
	pg, err := fileManager.LoadPage(title, filesDir, txtExtension)
	if err != nil {
		pg = &fileManager.Page{Title: title}
	}
	renderTemplate(responseWriter, pg, editTemplate)
}

// saveHandler: Get data from the request body and save it into a new file
func saveHandler(responseWriter http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	pg := fileManager.Page{
		Title: title,
		Body:  []byte(body),
	}
	err := pg.Save(filesDir, txtExtension, readWriteFileMode)
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

func healthCheckHandler(responseWriter http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	_ = r
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Header().Set("Content-Type", "application/json")

	_, err := io.WriteString(responseWriter, `{"alive": true}`)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

//renderTemplate: Load html template from the template cache and render page content to send back to client
func renderTemplate(responseWriter http.ResponseWriter, pg *fileManager.Page, templateName string) {
	err := templates.ExecuteTemplate(responseWriter, templateName, pg)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}
