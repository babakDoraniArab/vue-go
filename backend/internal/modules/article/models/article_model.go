package models

import (
	"github.com/babakDoraniArab/vue-go/internal/modules/user/models"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null"`
	Content string `gorm:"type:text;not null"`
	UserID  uint
	User    models.User
}
