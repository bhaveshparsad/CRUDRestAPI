package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllPerson(t *testing.T) {

	req, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(GetAllPerson)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	expected := `[{"id":1,"firstname":"Bhavesh","lastname":"Parsad","email":"Bhavesh@gmail.com"},{"firstname":"John","lastname":"Doe","email":"JohnDoe@gmail.com"}]`
	if r.Body.String() != expected {
		t.Errorf("Wrong Answer By Handler: got %v want %v",
			r.Body.String(), expected)
	}
}

func TestGetPersonByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(GetPersonByID)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	expected := `{"id":1,"firstname":"Bhavesh","lastname":"Parsad","email":"Bhavesh@gmail.com"}`
	if r.Body.String() != expected {
		t.Errorf("Wrong Answer By Handler: got %v want %v",
			r.Body.String(), expected)
	}
}

func TestCreatePerson(t *testing.T) {
	var jsonReq = []byte(`{"firstname":"Bhavesh","lastname":"Parsad","email":"Bhavesh@gmail.com"}`)

	req, err := http.NewRequest("POST", "/Create", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()

	control := http.HandlerFunc(CreatePerson)
	control.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}
	expected := `{"id":1,"firstname":"Bhavesh","lastname":"Parsad","email":"Bhavesh@gmail.com"}`
	if r.Body.String() != expected {
		t.Errorf("Wrong Answer By Handler: got %v want %v",
			r.Body.String(), expected)
	}
}

func TestUpdatePersonByID(t *testing.T) {

	var jsonReq = []byte(`{"id":2,"firstname":"John","lastname":"Doe","email":"JohnDoe@gmail.com"}`)

	req, err := http.NewRequest("PUT", "/update/{id}", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()
	control := http.HandlerFunc(UpdatePersonByID)
	control.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":2,"firstname":"John","lastname":"Doe","email":"JohnDoe@gmail.com"}`
	if r.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			r.Body.String(), expected)
	}
}

func TestDeletPersonByID(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/delete/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	control := http.HandlerFunc(DeletPersonByID)
	control.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("Wrong answer %v  got %v", status, http.StatusOK)
	}

	expected := `{"id":1,"firstname":"Bhavesh","lastname":"Parsad","email":"Bhavesh@gmail.com"}`
	if r.Body.String() != expected {
		t.Errorf("Wrong Answer By Handler: got %v want %v",
			r.Body.String(), expected)
	}

}