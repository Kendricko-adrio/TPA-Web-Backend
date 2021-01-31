package model

type Country struct {
	ID       int    `gorm:primaryKey`
	Name     string `json:"Name"`
	Code     string `json:"Code"`
}

