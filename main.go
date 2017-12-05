package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id, omitempty"`
	FirstName string   `json:"firstname, omitempty"`
	LastName  string   `json:"lastname, omitempty"`
	Address   *Address `json:"address, omitempty"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

var people []Person

func main() {

	router := mux.NewRouter()

	createData()

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}/{firstname}/{lastname}/{city}/{state}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createData() {
	people = append(people, Person{ID: "1", FirstName: "Rifan", LastName: "Ardiansyah", Address: &Address{City: "Sidoarjo",
		State: "Indonesia"}})

	people = append(people, Person{ID: "2", FirstName: "Putri", LastName: "Ardiansyah", Address: &Address{City: "Sidoarjo",
		State: "Indonesia"}})
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, person := range people {
		if person.ID == params["id"] {
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	person.FirstName = params["firstname"]
	person.LastName = params["lastname"]

	var address Address
	address.City = params["city"]
	address.State = params["state"]

	person.Address = &address

	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
