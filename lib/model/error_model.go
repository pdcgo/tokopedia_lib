package model

import "strings"

type Header struct {
	ProcessTime float64  `json:"processTime,omitempty"`
	Messages    []string `json:"messages"`
	Reason      string   `json:"reason"`
	ErrorCode   string   `json:"errorCode"`
	Typename    string   `json:"__typename"`
}

func (head *Header) Error() string {
	return strings.Join(head.Messages, "|")

}
