package authdto

type AuthRequest struct {
	Id       int    `json:"id"`
	Email    string `json:"email" binding:"required, email" gorm:"unique, not null" `
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullName" form:"fullName" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Role     string `json:"role" form:"role"`
}
type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
