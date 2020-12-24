package main

//import "encoding/json"

type Idea struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Idea  string `json:"idea"`
}
