package Models

type User struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
	JobTitle string `json:"jobtitle"`
}

var Users []User
