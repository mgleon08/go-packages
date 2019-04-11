package models

type User struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Info      `json:"info,omitempty"`
}
type Info struct {
	City  string `json:"city,omitempty"`
	Phone int    `json:"phone,omitempty"`
}
