package models

import "time"
const (
	UniqueConstraintUsername = "users_username_key"
	UniqueConstraintEmail    = "users_email_key"
  )
type User struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Name         string    `form:"name"`
	SurName      string    `form:"sur_name"`
	Birthday     time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	Email        string    `form:"email"`
	PasswordHash string    `form:"password_hash"`
	CreateTime   time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime     time.Time `form:"unixTime" time_format:"unix"`
}
