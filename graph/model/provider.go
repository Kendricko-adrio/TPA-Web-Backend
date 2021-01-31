package model

type Provider struct {
	ID   int    `gorm:primaryKey`
	Name string `json:"Name"`
}
