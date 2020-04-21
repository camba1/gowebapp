package fileManager

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// Page: Holds page we need to save/load
type Page struct {
	Title string
	Body  []byte
}

// fileSpec: File system details for a file
type fileSpec struct {
	Name     string
	ModTime  string
	Size     int64
	FullName string
}

//FileListing: List of filenames in a directory
type FileListing struct {
	Count     int
	Names     []string
	FullNames []string
	DirName   string
	Files     []fileSpec
}

//GetFileName: build file name base on tile + the file extension
func GetFileName(title, filesDir, txtExtension string) (string, error) {
	if title == "" || txtExtension == "" {
		return "", errors.New("invalid title or extension")
	}
	return filesDir + title + txtExtension, nil
}

// loadFileListing: Get the list of files in a given directory
func loadFileListing(dirName string) ([]os.FileInfo, error) {
	fileListing, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil, err
	}
	return fileListing, nil
}

//ListDirFiles: Load the count and names of the files in a directory into the FileListing struct
func ListDirFiles(directory, dateFormat string) (FileListing, error) {
	files := FileListing{DirName: directory}
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
		files = FileListing{
			Count:   len(fileSpecs),
			DirName: directory,
			Files:   fileSpecs,
		}
	}
	return files, err
}

//LoadPage: load page from disk into a page struct
func LoadPage(title string, filesDir, txtExtension string) (*Page, error) {
	fileName, err := GetFileName(title, filesDir, txtExtension)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{title, body}, nil
}

//save: Save a page to disk
func (pg *Page) Save(filesDir, txtExtension string, fileMode os.FileMode) error {
	fileName, err := GetFileName(pg.Title, filesDir, txtExtension)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, pg.Body, fileMode)

}

//saveDisplayTestPage: Save a string into a file. Reload the file and display content.
func SaveDisplayTestPage(filesDir, txtExtension string, fileMode os.FileMode) {
	pageTitle := "History"
	pg := Page{
		Title: pageTitle,
		Body:  []byte("This is history!"),
	}
	err := pg.Save(filesDir, txtExtension, fileMode)
	if err != nil {
		panic(err)
	}
	pgLd, err := LoadPage(pageTitle, filesDir, txtExtension)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(pgLd.Title + ": " + string(pgLd.Body))
	}
}
