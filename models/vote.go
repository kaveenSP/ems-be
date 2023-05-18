package models

type Vote struct {
	VoteId    string   `json:"voteId"`
	Title     string   `json:"title"`
	TeacherId string   `json:"teacherId"`
	Options   []string `json:"options"`
	Status    string   `json:"status"`
}
