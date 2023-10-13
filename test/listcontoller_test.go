package test 

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlePostRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "/add", strings.NewReader("todolist=test&deadline=2023-12-31"))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Call the handler function, passing in the ResponseRecorder and Request.
	TestHandlePostRequest()

	// Check the status code is what you expect.
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusSeeOther)
	}

	// You can further check the response body if needed.
	// responseBody := rr.Body.String()
	// Add assertions for the response body here.
}