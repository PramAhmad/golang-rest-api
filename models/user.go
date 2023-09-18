package models

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"nama"`
	Kelas string `json:"kelas"`
}
