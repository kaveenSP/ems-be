package models

type Teacher struct {
	TeacherId     string `json:"teacherId"`
	Role          string `json:"role"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	ContactNumber string `json:"contactNumber"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
	DOB           string `json:"dob"`
	Address       string `json:"address"`
	NIC           string `json:"nic"`
	ImagePath     string `json:"imagePath"`
	Password      string `json:"password"`
}
