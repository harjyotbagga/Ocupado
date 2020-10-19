package models

type Error struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
}
