package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"tc_go/controllers"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/users", controllers.GetPersons).Methods("GET")
	r.HandleFunc(`/user/{id:\d+}`, controllers.GetPersonById).Methods("GET")
	r.HandleFunc("/cars", controllers.GetCars).Methods("GET")
	r.HandleFunc(`/car/{id:\d+}`, controllers.GetCarById).Methods("GET")
	r.HandleFunc(`/user/{id:\d+}/cars`, controllers.GetPersonsCarsById).Methods("GET")
	r.HandleFunc(`/user/{id:\d+}/with/cars`, controllers.GetPersonWithCarById).Methods("GET")
	r.HandleFunc(`/houses`, controllers.GetHousesWithParkingPlacesAndPersons).Methods("GET")
	r.HandleFunc(`/house/{id:\d+}`, controllers.GetHouseWithParkingPlacesAndPersonsById).Methods("GET")

	//r.Use(app.JwtAuthentication) //attach JWT auth middleware

	//r.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, r) //Launch the app, visit localhost:8080
	if err != nil {
		fmt.Print(err)
	}
}
