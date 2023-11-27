package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type User struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

var users []User

func main() {
    router := mux.NewRouter()

    // Route to get all users
    router.HandleFunc("/users", GetUsers).Methods("GET")

    // Route to get a single user by ID
    router.HandleFunc("/users/{id}", GetUser).Methods("GET")

    // Route to create a new user
    router.HandleFunc("/users", CreateUser).Methods("POST")

    // Route to update user information
    router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")

    // Route to delete a user
    router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", router))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r) // Get params
    for _, item := range users {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&User{})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user User
    _ = json.NewDecoder(r.Body).Decode(&user)
    users = append(users, user)
    json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range users {
        if item.ID == params["id"] {
            users = append(users[:index], users[index+1:]...)
            var user User
            _ = json.NewDecoder(r.Body).Decode(&user)
            user.ID = params["id"]
            users = append(users, user)
            json.NewEncoder(w).Encode(user)
            return
        }
    }
    json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range users {
        if item.ID == params["id"] {
            users = append(users[:index], users[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(users)
}
