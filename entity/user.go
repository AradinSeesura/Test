package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	StudentID string `valid:"required~StudentID is required, matches(^[BMD]\\d{7}$)"`
	Email     string `valid:"required~Email is required, email~Email is invalid"`
	Phone     string `valid:"required~Phone is required, stringlength(10|10)"`
	Profile   string `gorm:"type:longtext"`
	LinkedIn  string `valid:"required~required is required, url~Url is invalid"`
	Password  string `valid:"required~Password is required, minstringlength(8)~Password must be at leasts 8"`
	Age       int    `valid:"required~Age is required, range(18|99)~Age must be between 18 and 99"`
}