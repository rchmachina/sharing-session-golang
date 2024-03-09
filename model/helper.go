package model

type Response struct {
	Message interface{} `json:"message"`
	Status  string      `json:"status"`
}
