package models

import (
	"errors"
	"html"
	"log"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	UniqueConstraintUsername = "users_username_key"
	UniqueConstraintEmail    = "users_email_key"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"size:255;not null;unique" json:"first_name" form:"first_name"`
	LastName  string `gorm:"size:255;not null;unique" json:"last_name" form:"last_name"`
	Email     string `gorm:"size:100;not null;unique" json:"email" form:"email"`
	Age       uint8  `json:"age" form:"age"`
	Password  string `json:"password" form:"password"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
func (u *User) Prepare() {
	u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
}
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.FirstName == "" {
			return errors.New("Required Firstname")
		}
		if u.LastName == "" {
			return errors.New("Required Lastname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.FirstName == "" {
			return errors.New("Required Firstname")
		}
		if u.LastName == "" {
			return errors.New("Required Lastname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	}

}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	u.BeforeSave()
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	err = gorm.ErrRecordNotFound
	if err != nil {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"first_name": u.FirstName,
			"last_name":  u.LastName,
			"email":      u.Email,
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
