package main

type Event struct {
	Name string `json:"what is your name?"`
	Age  int    `json:"how old are you"`
}

type Response struct {
	Message string `json:"answer"`
}
