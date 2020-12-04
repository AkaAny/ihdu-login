package main

type LoginResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"msg"`
	UserName string `json:"userName"`
}
