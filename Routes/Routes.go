package Routes

import (
	"github.com/akshat-sahijpal/Crud-Go/Controllers"
	"github.com/gorilla/mux"
)

var RegisterUser = func(router *mux.Router) {
	router.HandleFunc("/register", Controllers.RegisterUser).Methods("POST")
}
var GetAllUser = func(router *mux.Router) {
	router.HandleFunc("/getAllUsers", Controllers.GetAllUsers).Methods("GET")
}
var GetUserForUserId = func(router *mux.Router) {
	router.HandleFunc("/getUser/{id}", Controllers.GetUserForUserId).Methods("GET")
}

var DeleteUserForId = func(router *mux.Router) {
	router.HandleFunc("/deleteUser/{id}", Controllers.DeleteUserForId).Methods("DELETE")
}
var UpdateUserForId = func(router *mux.Router) {
	router.HandleFunc("/updateUser/{id}", Controllers.UpdateUserForId).Methods("PUT")
}
