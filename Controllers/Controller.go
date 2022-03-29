package Controllers

import (
	"encoding/json"
	"github.com/akshat-sahijpal/Crud-Go/Models"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
)

func RegisterUser(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var newUser Models.User
	_ = json.NewDecoder(req.Body).Decode(&newUser)
	newUser.Id = string(rand.Intn(100))
	Models.Users = append(Models.Users, newUser)
	json.NewEncoder(writer).Encode(newUser)
}
func GetAllUsers(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(Models.Users); err != nil {
		return
	}
}
func GetUserForUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	print(&params)
	for _, v := range Models.Users {
		if v.Id == params["id"] {
			print(v.Id)
			json.NewEncoder(w).Encode(v)
			break
		}
	}
}
func DeleteUserForId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // returns arg parameters
	for ind, v := range Models.Users {
		if v.Id == params["id"] {
			Models.Users = append(Models.Users[:ind], Models.Users[ind+1:]...)
			break
		}
	}
}
func UpdateUserForId(w http.ResponseWriter, r *http.Request) {
	// Delete the user than add a new user
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // returns arg parameters
	for ind, v := range Models.Users {
		if v.Id == params["id"] {
			Models.Users = append(Models.Users[:ind], Models.Users[ind+1:]...)
			var newUser Models.User
			_ = json.NewDecoder(r.Body).Decode(&Models.Users)
			newUser.Id = string(rune(rand.Intn(100)))
			Models.Users = append(Models.Users, newUser)
			return
		}
	}
}
