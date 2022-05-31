package models

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Deptid   string `json:"deptid"`
	Roles    string `json:"roles"`
}
