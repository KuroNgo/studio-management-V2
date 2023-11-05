package model

type HTTPError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
