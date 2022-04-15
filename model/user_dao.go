package model

// User 用户关键数据
type User struct {
	Uid         string `gorm:"primaryKey"`
	StudentID   string `gorm:"unique"`
	RealName    string
	PhoneNumber string `gorm:"unique"`
	XcxOpenID   string `gorm:"unique;column:xcx_openid"`
	AppOpenID   string `gorm:"unique;column:app_openid"`
	UnionID     string `gorm:"unique;column:union_id"`
}

type Student struct {
	StudentID string `gorm:"primaryKey;column:student_id"`
	RealName  string
}
