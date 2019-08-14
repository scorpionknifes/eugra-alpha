package main

//Status model for REST API response
type Status struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
