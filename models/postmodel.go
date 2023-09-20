package models

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	Name string
	Type string
}
