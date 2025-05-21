package models

type Error struct {
	ErrorName  string `json:"error_name"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
