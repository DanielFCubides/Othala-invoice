package config

import "gorm.io/gorm"

type Connection interface {
	GetDatabase() *gorm.DB
}
