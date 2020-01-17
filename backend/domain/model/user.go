package model

import "time"

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	UUID      string     `json:"uuid"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// TableName gets table name
func (u *User) TableName() string {
	return "users"
}

