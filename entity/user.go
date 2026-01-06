package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// แก้: เพิ่มชื่อ Field ใน error message เพื่อความชัดเจน
	StudentID string `valid:"required~StudentID is required, matches(^[BMD]\\d{7}$)"`

	// แก้: ลบเครื่องหมาย " ที่คั่นกลางออก (Tag ต้องเป็น string ยาวต่อเนื่องกัน)
	Email string `valid:"required~Email is required, email~Email is invalid"`

	Phone string `valid:"required~Phone is required, stringlength(10|10)"`

	Profile string `gorm:"type:longtext"`

	// แก้: เติม ~ หลัง required และแก้รูปแบบข้อความ
	LinkedIn string `valid:"required~LinkedIn is required, url~Url is invalid"`

	// แก้: แก้คำผิด at lasts -> at least
	Password string `valid:"required~Password is required, minstringlength(8)~Password must be at least 8 characters"`

	// แก้: แก้คำผิด rage -> range
	Age int `valid:"required~Age is required, range(18|99)~Age must be between 18 and 99"`
}