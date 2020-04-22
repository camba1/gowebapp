package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_healthCheckHandler(t *testing.T) {
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

func Test_viewHandler(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		wantErr  bool
		expected string
	}{
		{name: "View test history page", title: "History", wantErr: false, expected: ""},
		{name: "Invalid page name", title: "xyz", wantErr: true, expected: "<a href=\"/edit/xyz\">Found</a>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			title := tt.title
			req, err := http.NewRequest("GET", "/view/"+title, nil)
			if err != nil {
				t.Fatal(err)
			}
			recorder := httptest.NewRecorder()
			handler := makeHandler(viewHandler)
			handler.ServeHTTP(recorder, req)

			if tt.wantErr {
				if status := recorder.Code; status != http.StatusFound {
					t.Errorf("Attempting to view non existing file returned the wrong http code: got %v, want %v", status, http.StatusFound)
				}
				if !strings.Contains(recorder.Body.String(), "/edit/"+title) {
					t.Errorf("Expected redirection. Wanted address /edit/"+title+", got %v ", recorder.Body.String())
				}
			} else {
				if status := recorder.Code; status != http.StatusOK {
					t.Errorf("Attempting to view "+title+" file returned the wrong http code: got %v, want %v", status, http.StatusOK)
				}

				if recorder.Body.String() == "" {
					t.Errorf("Recieved an empty body %v", recorder.Body.String())
				}
			}

		})
	}
}
