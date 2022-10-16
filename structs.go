package main

type Response struct {
	M Magic `json:"magic"`
}
type Magic struct {
	Q string `json:"question"`
	A string `json:"answer"`
	T string `json:"type"`
}
