package user

// type Test struct {
// 	Apple  string `header:"apple" `
// 	Banana string `json:"banana,omitempty" binding:"required" gorm:"column:banana" `
// }

type ReqUserSignup struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required" `
}

type ReqUserLogout struct {
	UserId string `json:"user_id" binding:"required"`
}
