package model

import (
	"be-blog/pkg/idgen"
	"gorm.io/gorm"
)

type Role struct {
	ID   string `gorm:"primaryKey;size:26"`
	Name string `gorm:"uniqueIndex;not null"`
	TimeStamps
}

func (u *Role) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = idgen.NewULID()
	return
}
