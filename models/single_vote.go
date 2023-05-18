package models

type SingleVote struct {
	VoteId    string `json:"voteId"`
	StudentId string `json:"studentId"`
	Option    string `json:"option"`
}
