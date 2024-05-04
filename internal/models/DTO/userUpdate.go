package DTO

type UserUpdate struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Telegram *string `json:"telegram"`
	Password *string `json:"password"`
	Status   *string `json:"status"`
}
