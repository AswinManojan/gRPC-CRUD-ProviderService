package model

import "gorm.io/gorm"

type Provider struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique; not null"`
}


type ProviderLogin struct {
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"unique;not null"`
}
