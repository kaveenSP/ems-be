package models

type Admin struct {
	AdminId   string `json:"adminId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	DOB       string `json:"dob"`
	Password  string `json:"password"`
}
