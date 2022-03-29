package Controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/register", writer.Body)
	RegisterUser(writer, request)

	if writer.Code != http.StatusCreated { // 201 is returned when a post request is successful
		t.Errorf("Error creating a new user")
	}

}
func TestGetAllUsers(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/getAllUsers", nil)

	GetAllUsers(writer, request)
	if writer.Code == http.StatusNoContent { // No Content
		t.Errorf("No data displayed")
	}
}

func TestGetUserForUserId(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/getUser/{id}", nil)

	GetUserForUserId(writer, request)
	if writer.Code == http.StatusNoContent { // No Content
		t.Errorf("No data displayed")
	}
}
func TestDeleteUserForId(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/deleteUser/{id}", nil)
	DeleteUserForId(writer, request)
	if writer.Code != http.StatusAccepted { // No Content
		t.Errorf("Error Deleting")
	}
}
func TestUpdateUserForId(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPut, "/updateUser/{id}", nil)
	UpdateUserForId(writer, request)
	if writer.Code != http.StatusCreated {
		t.Errorf("Error updating given user")
	}
}
