package model

type Error struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Typename    string `json:"__typename"`
}

type ErrorMsg struct {
	Message  string `json:"message"`
	Typename string `json:"__typename"`
}

type Fallback struct {
	Message  string `json:"message"`
	HTML     string `json:"html"`
	Typename string `json:"__typename"`
}
