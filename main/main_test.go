package main

import (
	"CRUDRestAPI/controllers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateContact(t *testing.T) {

	newContact := []byte(`
		{
			"FirstName" : "first",
			"LastName"	: "last",
			"Email" 	: ""
		}
	`)

	//creating request
	req, err := http.NewRequest("POST", "localhost:12345", strings.NewReader(string(newContact)))
	if err != nil {
		t.Fatal("Could not create request for CreateContact.\n", err)
	}
	//Recorder or ResponseWriter
	rec := httptest.NewRecorder()

	//passing test database using context

	service.CreateContact(rec, req)

	response := rec.Result()

	if response.StatusCode != http.StatusOK {

		t.Error("Expected Status Ok, got ", response.StatusCode)
	} else {
		t.Log(`"
	CreateContact Passed:
	Status Code: "`, response.StatusCode)
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Could not read response body")
	}

	t.Log("\nTest Result : ", string(result))

}

func TestGetAllContacts(t *testing.T) {
	
	newContact := []byte(`
		{
			"FirstName" : "first",
			"LastName"	: "last",
			"Email" 	: "firstlast@gmail.com"
		}
	`)

	//creating request
	req, err := http.NewRequest("POST", "localhost:12345", strings.NewReader(string(newContact)))
	if err != nil {
		t.Fatal("Could not create request for CreateContact.\n", err)
	}
	//Recorder or ResponseWriter
	rec := httptest.NewRecorder()

	//passing test database using context

	service.CreateContact(rec, req)

	//creating request
	req, err = http.NewRequest("GET", "localhost:12345", nil)
	if err != nil {
		t.Fatal("Could not create request for GetAllContacts.\n", err)
	}
	//Recorder or ResponseWriter
	rec = httptest.NewRecorder()

	service.GetAllContacts(rec, req)

	response := rec.Result()

	if response.StatusCode != http.StatusOK {

		t.Error("Expected Status Ok, got ", response.StatusCode)
	} else {
		t.Log(`"
		GetAllContacts Passed:
		Status Code: "`, response.StatusCode)
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Could not read response body")
	}
	t.Log("\nTest Result : ", string(result))
}

func TestUpdateContact(t *testing.T) {
	
	newContact := []byte(`
		{
			"FirstName" : "first",
			"LastName"	: "last",
			"Email" 	: ""
		}
	`)

	//creating request
	req, err := http.NewRequest("POST", "localhost:12345", strings.NewReader(string(newContact)))
	if err != nil {
		t.Fatal("Could not create request for CreateContact.\n", err)
	}
	//Recorder or ResponseWriter
	rec := httptest.NewRecorder()

	//passing test database using context

	service.CreateContact(rec, req)
	updateContact := []byte(`
		{
			"FirstName" : "first1",
			"LastName"	: "last1",
			"Email" 	: "first1last1@gmail.com"
			
		}
	`)

	req, err = http.NewRequest("PUT", "localhost:12345", strings.NewReader(string(updateContact)))
	if err != nil {
		t.Fatal("Could not create request for UpdateContact.\n", err)
	}
	//Recorder or ResponseWriter
	rec = httptest.NewRecorder()

	//fake gorilla/mux var
	vars := map[string]string{}

	// setting var to the req
	req = mux.SetURLVars(req, vars)

	service.UpdateContact(rec, req)

	response := rec.Result()

	if response.StatusCode != http.StatusOK {

		t.Error("Expected Status Ok, got ", response.StatusCode)
	} else {
		t.Log(`"
	UpdateContact Passed:
	Status Code: "`, response.StatusCode)
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Could not read response body")
	}

	t.Log("\nTest Result : ", string(result))

}

func TestDeleteContact(t *testing.T) {
	
	newContact := []byte(`
		{
			"FirstName" : "first",
			"LastName"	: "last",
			"Email" 	: ""
		}
	`)

	//creating request
	req, err := http.NewRequest("POST", "localhost:12345", strings.NewReader(string(newContact)))
	if err != nil {
		t.Fatal("Could not create request for CreateContact.\n", err)
	}
	//Recorder or ResponseWriter
	rec := httptest.NewRecorder()

	//passing test database using context

	service.CreateContact(rec, req)

	req, err = http.NewRequest("DELETE", "localhost:12345", nil)
	if err != nil {
		t.Fatal("Could not create request for DeleteContact.\n", err)
	}
	//Recorder or ResponseWriter
	rec = httptest.NewRecorder()

	service.DeleteContact(rec, req)

	response := rec.Result()

	if response.StatusCode != http.StatusOK {

		t.Error("Expected Status Ok, got ", response.StatusCode)
	} else {
		t.Log(`"
	DeleteContact Passed:
	Status Code: "`, response.StatusCode)
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Could not read response body")
	}

	t.Log("\nTest Result : ", string(result))

}