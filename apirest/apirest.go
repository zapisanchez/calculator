package apirest

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"ticketCalc/users"

	"github.com/gorilla/mux"
)

/*Users : List of Users to get in Memory */
var Users map[int]users.User

/*CreateRouter : Create the router with their CRUD calls*/
func CreateRouter() {
	Users = make(map[int]users.User)
	log.Println("Default users: ", Users)
	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUser).Methods("POST")
	//router.HandleFunc("/users", GetUsers).Methods("GET")
	//router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	//router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

/*CreateUser : Create an User and store it. POST Call */
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user users.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if _, ok := Users[user.ID]; ok {
		http.Error(w, "Already exists an user with the same id", http.StatusConflict)
		return
	}

	Users[user.ID] = user

	log.Println("User: ", user, "added")
}

/*GetUsers : return a list of users. GET Call */
func GetUsers(w http.ResponseWriter, r *http.Request) {

	userList := make([]users.User, 0)

	for _, value := range Users {

		userList = append(userList, value)

	}

	json.NewEncoder(w).Encode(userList)
}

/*GetUser : return a given user, if exist.
if id is not the correct, return err 400
if user not exist, return 404 err
*/
func GetUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if user, ok := Users[id]; ok {
		json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

/*DeleteUser : Delete an user from the map, if its given.
Otherwise, return 400 or 404 error*/
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if user, ok := Users[id]; ok {

		delete(Users, id)
		log.Println("User: ", user, "removed")

	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}

}
