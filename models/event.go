package models

type Event struct {
	EventId     string `json:"eventId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	StartTime   string `json:"startTime"`
	EndDate     string `json:"endDate"`
	EndTime     string `json:"endTime"`
	Location    string `json:"location"`
	Incharge    string `json:"incharge"`
	ImagePath   string `json:"imagePath"`
	Open        bool   `json:"open"`
}
