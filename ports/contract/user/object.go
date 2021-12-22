package user

import "time"

type Object struct {
	Id        string
	Name      string
	Role      string
	Address   string
	CreatedAt time.Time
}
