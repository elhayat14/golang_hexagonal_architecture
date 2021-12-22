package user

import "time"

type CreateUserRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}
type UpdateUserRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
type DefaultResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Role      string    `bson:"role"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}
