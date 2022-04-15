package model

// UserJson 接收用户关键数据
type UserJson struct {
	StudentID   string `json:"student_id,omitempty"`
	RealName    string `json:"real_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Code        string `json:"code,omitempty"`
}
