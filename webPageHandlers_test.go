package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//func Test_editHandler(t *testing.T) {
//	type args struct {
//		responseWriter http.ResponseWriter
//		r              *http.Request
//		title          string
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}
//
//func Test_handler(t *testing.T) {
//	type args struct {
//		responseWriter http.ResponseWriter
//		r              *http.Request
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}

func Test_healthCheckHandler(t *testing.T) {
	//type args struct {
	//	responseWriter http.ResponseWriter
	//	r              *http.Request
	//}
	tests := []struct {
		name string
		//args args
	}{
		{name: "Check health check"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := `{"alive": true}`
			req, err := http.NewRequest("GET", "/healthCheck", nil)
			if err != nil {
				t.Fatal(err)
			}
			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(healthCheckHandler)
			handler.ServeHTTP(recorder, req)

			if status := recorder.Code; status != http.StatusOK {
				t.Errorf("Health Check returned the wrong http code: got %v, want %v", status, http.StatusOK)
			}

			if recorder.Body.String() != expected {
				t.Errorf("Health Check bodey contains %v, want %v", recorder.Body.String(), expected)
			}
		})
	}
}

//func Test_makeHandler(t *testing.T) {
//	type args struct {
//		fn func(http.ResponseWriter, *http.Request, string)
//	}
//	tests := []struct {
//		name string
//		args args
//		want http.HandlerFunc
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := makeHandler(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("makeHandler() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_newHandler(t *testing.T) {
//	type args struct {
//		responseWriter http.ResponseWriter
//		r              *http.Request
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}
//
//func Test_renderTemplate(t *testing.T) {
//	type args struct {
//		responseWriter http.ResponseWriter
//		pg             *fileManager.Page
//		templateName   string
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}
//
//func Test_saveHandler(t *testing.T) {
//	type args struct {
//		responseWriter http.ResponseWriter
//		r              *http.Request
//		title          string
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}
//
//func Test_startServer(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}
//
//func Test_viewHandler(t *testing.T) {
//	type args struct {
//		responseWriter http.ResponseWriter
//		r              *http.Request
//		title          string
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}
