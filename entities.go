package main

import "time"

type User struct {
	ID       int    `json:"id"`
	CI       string `json:"ci"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Sex      string `json:"sex"`
	Ege      string `json:"ege"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Created  time.Time
}
