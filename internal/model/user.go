package model

type User struct {
	ID           uint   `gorm:"primaryKey"`
	FullName     string `gorm:"size:45;not null"`
	Username     string `gorm:"uniqueIndex;not null"`
	Email        string `gorm:"uniqueIndex,not null"`
	Password     string
	TokenVersion string
	Roles        []*Role `gorm:"many2many:user_roles;"`
	TimeStamps
}
