package entity

type App struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Secret string `json:"secret" db:"secret"`
}
