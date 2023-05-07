package models

type Student struct {
	StudentId     string `json:"studentId"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	ContactNumber string `json:"contactNumber"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
	DOB           string `json:"dob"`
	Address       string `json:"address"`
	Password      string `json:"password"`
}
