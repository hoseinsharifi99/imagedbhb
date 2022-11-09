package model

import "github.com/jinzhu/gorm"

type GameImage struct {
	gorm.Model
	GameID       uint   `json:"userId"`
	ImageAddress string `json:"image"`
}
