package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"io"
	"os"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	jwt "github.com/dgrijalva/jwt-go"
)

type Location struct {
	Id int `json:"Id"`
    City string `json:"City"`
    Address string `json:"Address"`
    PostCode string `json:"PostCode"`
}

type User struct {
	Id int `json:"Id"`
    Username string `json:"Username"`
    PreferredLocation string `json:"PreferredLocation"`
}

var mySigningKey = []byte(os.Getenv("TOKEN_SECRET"))
var Locations []Location
var Users []User

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func createLocation(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)
    var location Location
    err := json.Unmarshal(reqBody, &location)
    if(err != nil) {
        fmt.Fprintf(w, err.Error())
    }

	index := len(Locations) - 1
	location.Id = Locations[index].Id + 1
    Locations = append(Locations, location)

    json.NewEncoder(w).Encode(location)
}

func updatePreferredLocation (w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)
    var userObj User
    json.Unmarshal(reqBody, &userObj)
	i := 0
    for _, u := range Users {
        if u.Username == userObj.Username {
            fmt.Println("Match")
			Users[i].PreferredLocation = userObj.PreferredLocation
			u.PreferredLocation = userObj.PreferredLocation
			fmt.Println(u)
            json.NewEncoder(w).Encode(u)
        }
		i++
    }
}

func returnAllLocations(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllLocations")
    json.NewEncoder(w).Encode(Locations)
}


func returnAllUsers(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllUsers")
    json.NewEncoder(w).Encode(Users)
}

func getLocation(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key, _ :=  strconv.Atoi(vars["id"])

    for _, loc := range Locations {
        var id int = loc.Id
        if id == key {
            json.NewEncoder(w).Encode(loc)
        }
    }
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")

    io.WriteString(w, `{"alive": true}`)
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        if r.Header["Token"] != nil {
            token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token)(interface{}, error){
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
                    return nil, fmt.Errorf("There was an error")
                }
                return mySigningKey, nil
            })

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            if token.Valid {
                endpoint(w, r)
            }
        } else {
            fmt.Fprintf(w, "Not Authorized")
        }
    })
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homePage)
	router.Handle("/api/metrics", promhttp.Handler())
	router.HandleFunc("/api/health", healthCheck).Methods("GET")

	router.Handle("/api/locations", isAuthorized(returnAllLocations)).Methods("GET")
	router.Handle("/api/location", isAuthorized(createLocation)).Methods("POST")
	router.Handle("/api/location/{id}", isAuthorized(getLocation)).Methods("GET")

	router.Handle("/api/users", isAuthorized(returnAllUsers)).Methods("GET")
	router.Handle("/api/user/update", isAuthorized(updatePreferredLocation)).Methods("PUT")

    header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Token"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
    origins := handlers.AllowedOrigins([]string{"*"})

    log.Fatal(http.ListenAndServe(":8081", handlers.CORS(header, methods, origins)(router)))
}

func main() {
	Locations = []Location{
        Location{Id: 1, City: "Singapore", Address: "9 Battery Road", PostCode: "049910"},
        Location{Id: 2, City: "Singapore", Address: "380 Jalan Besar", PostCode: "209000"},
    }
	Users = []User{
        User{Id: 1, Username: "Andy", PreferredLocation: "9 Battery Road" },
        User{Id: 2, Username: "Sean", PreferredLocation: "9 Battery Road" },
    }
    handleRequests()
}