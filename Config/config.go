package Config

import (
	"github.com/akshat-sahijpal/Crud-Go/Models"
	"github.com/akshat-sahijpal/Crud-Go/Routes"
	"github.com/gorilla/mux"
	"net/http"
)

func ConfigureBasicSettings() {
	router := mux.NewRouter()
	Models.Users = append(Models.Users, Models.User{Id: "1", UserName: "@mount.akshat", FullName: "Akshat Sahijpal", Age: 21, JobTitle: "SDE"})
	Models.Users = append(Models.Users, Models.User{Id: "2", UserName: "@brown.201", FullName: "Minner", Age: 24, JobTitle: "SDE3"})
	Routes.RegisterUser(router)
	Routes.GetAllUser(router)
	Routes.GetUserForUserId(router)
	Routes.DeleteUserForId(router)
	Routes.UpdateUserForId(router)
	if err := http.ListenAndServe(":8000", router); err != nil {
		return
	}
}
