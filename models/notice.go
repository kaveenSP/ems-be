package models

type Notice struct {
	NoticeId string `json:"noticeId"`
	Subject  string `json:"subject"`
	Message  string `json:"message"`
}
