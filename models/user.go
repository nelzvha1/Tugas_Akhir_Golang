package models

type User struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Kelas   string `json:"kelas"`
	Jurusan string `json:"jurusan"`
	Email   string `json:"email" gorm:"unique"`
}
