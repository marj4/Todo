package models

type Tasks struct {
	ID          int    `json:"id"`
	Nametask    string `json:"nametask"`
	Discription string `json:"discription"`
	Status      bool   `json:"status"`
}
