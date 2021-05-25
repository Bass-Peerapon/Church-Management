package models

import "time"

type (
	Login struct {
		User     string `form:"user" json:"user" xml:"user"  binding:"required"`
		Password string `form:"password" json:"password" xml:"password" binding:"required"`
	}

	User struct {
		Name       string    `form:"name"`
		SurName    string    `form:"sur_name"`
		Address    string    `form:"address"`
		Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
		CreateTime time.Time `form:"createTime" time_format:"unixNano"`
		UnixTime   time.Time `form:"unixTime" time_format:"unix"`
	}
)
