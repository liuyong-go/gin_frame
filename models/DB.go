package models

type Admin struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ShortUrlMap struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Lurl string `json:"lurl"`
	Surl string `json:"surl"`
}
