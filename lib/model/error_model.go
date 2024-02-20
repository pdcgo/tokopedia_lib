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

func (head *Header) IsProductFull() bool {
	for _, message := range head.Messages {
		if strings.Contains(message, "Jumlah produk yang dapat Anda tambahkan dibatasi") {
			return true
		}

		if strings.Contains(message, "sudah melewati batas kuota") {
			return true
		}
	}

	return false
}

func (head *Header) IsBanned() bool {
	for _, message := range head.Messages {
		if strings.Contains(message, "toko ditangguhkan permanen") {
			return true
		}
	}

	return false
}

func (head *Header) GetBannedWord() string {
	for _, message := range head.Messages {
		if strings.Contains(message, "Nama Produk memuat kata") {
			return message
		}

		if strings.Contains(message, "Deskripsi memuat kata") {
			return message
		}
	}

	return ""
}
