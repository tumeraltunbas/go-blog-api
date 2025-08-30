package entities

import "time"

type User struct {
	Id        string
	Email     string
	Password  string
	IsActive  bool
	CreatedAt time.Time
}
