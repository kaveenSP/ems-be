package models

type RegisteredEvent struct {
	StudentId   string `json:"studentId"`
	StudentName string `json:"studentName"`
	EventId     string `json:"eventId"`
	EventName   string `json:"eventName"`
	Attendance  bool   `json:"attendance"`
}
