package main

import (
	//"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	//"io/ioutil"
)

func TestGetUsers(t *testing.T){
	prepareData()
	req, _ := http.NewRequest("GET", "/api/users", nil)
	response := executeRequest(req)
	
	checkResponseCode(t, http.StatusOK, response.Code)
}

func prepareData(){
	Locations = []Location{
        Location{Id: 1, City: "Singapore", Address: "9 Battery Road", PostCode: "049910"},
    }
	Users = []User{
        User{Id: 1, Username: "Andy", PreferredLocation: "9 Battery Road" },
    }
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
	main := http.HandlerFunc(returnAllUsers)
    main.ServeHTTP(rr, req)

    return rr
}