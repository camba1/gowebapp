package fileManager

import (
	"fmt"
	"os"
	"testing"
)

func TestGetFileName(t *testing.T) {
	type args struct {
		title        string
		filesDir     string
		txtExtension string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Positive test",
			args:    args{title: "MyTitle", filesDir: "files/", txtExtension: ".txt"},
			want:    "files/MyTitle.txt",
			wantErr: false,
		},
		{name: "No directory provided",
			args:    args{title: "MyTitle", filesDir: "", txtExtension: ".txt"},
			want:    "MyTitle.txt",
			wantErr: false,
		},
		{name: "Missing title",
			args: args{title: "", filesDir: "files/", txtExtension: ".txt"},
			//want: "invalid title or extension",
			wantErr: true,
		},
		{name: "Missing extension",
			args: args{title: "MyTitle", filesDir: "files/", txtExtension: ""},
			//want: "invalid title or extension",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFileName(tt.args.title, tt.args.filesDir, tt.args.txtExtension)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFileName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListDirFiles(t *testing.T) {
	type args struct {
		directory  string
		dateFormat string
	}
	tests := []struct {
		name    string
		args    args
		want    FileListing
		wantErr bool
	}{
		{name: "Find History file",
			args: args{directory: "../files/", dateFormat: "Jan-02-06"},
			want: FileListing{
				Count:     0,
				Names:     nil,
				FullNames: nil,
				DirName:   "../files/",
				Files:     []fileSpec{{Name: "History", FullName: "History.txt"}},
			},
			wantErr: false,
		},
		{name: "Invalid file",
			args:    args{directory: "xyz/", dateFormat: "Jan-02-06"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListDirFiles(tt.args.directory, tt.args.dateFormat)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListDirFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				foundHistory := false
				for _, file := range got.Files {
					fmt.Println(file.FullName + " " + file.Name + " - " + tt.want.Files[0].FullName + " " + tt.want.Files[0].Name)
					if file.FullName != tt.want.Files[0].FullName || file.Name != tt.want.Files[0].Name {
						foundHistory = true
					}
				}
				if !foundHistory {
					t.Errorf("ListDirFiles() got = %v, missing %s", got, tt.want.Files[0].FullName)
				}
			}
		})
	}
}

func TestPage_Save(t *testing.T) {
	type fields struct {
		Title string
		Body  []byte
	}
	type args struct {
		filesDir     string
		txtExtension string
		fileMode     os.FileMode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Save History File",
			fields: fields{
				Title: "History",
				Body:  []byte("This is history!"),
			},
			args: args{
				filesDir:     "../files/",
				txtExtension: ".txt",
				fileMode:     0600,
			},
			wantErr: false,
		},
		{name: "Missing file name",
			fields: fields{
				Title: "",
				Body:  []byte("This is history!"),
			},
			args: args{
				filesDir:     "files",
				txtExtension: ".txt",
				fileMode:     0600,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pg := &Page{
				Title: tt.fields.Title,
				Body:  tt.fields.Body,
			}
			if err := pg.Save(tt.args.filesDir, tt.args.txtExtension, tt.args.fileMode); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSaveDisplayTestPage(t *testing.T) {
	type args struct {
		filesDir     string
		txtExtension string
		fileMode     os.FileMode
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Save and Display test page",
			args: args{
				filesDir:     "files",
				txtExtension: ".txt",
				fileMode:     0600,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
