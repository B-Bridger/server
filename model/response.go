// model/response.go
package model

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type OKResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

type TokenResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	User    User   `json:"user"`
	Token   string `json:"token"`
}
