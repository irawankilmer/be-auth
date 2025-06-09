package model

type Role struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `gorm:"uniqueIndex;not null"`
	Users []*User `gorm:"many2many:user_roles;"`
	TimeStamps
}
