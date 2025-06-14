package model

import (
	"github.com/irawankilmer/be-auth/pkg/idgen"
	"gorm.io/gorm"
)

type User struct {
	ID           string `gorm:"primaryKey;size:26"`
	FullName     string `gorm:"size:45;not null"`
	Username     string `gorm:"uniqueIndex;not null"`
	Email        string `gorm:"uniqueIndex,not null"`
	Password     string
	TokenVersion string
	Roles        []*Role `gorm:"many2many:user_roles;"`
	TimeStamps
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = idgen.NewULID()
	return
}
