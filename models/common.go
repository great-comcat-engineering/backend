package models

type ApiResponse[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Content T      `json:"content"`
}
