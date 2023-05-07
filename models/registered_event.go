package models

type RegisteredEvent struct {
	StudentId  string `json:"studentId"`
	EventId    string `json:"eventId"`
	Attendance bool   `json:"attendance"`
}
