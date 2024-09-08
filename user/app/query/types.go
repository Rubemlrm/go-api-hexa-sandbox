package query

import "time"

type ID int

type User struct {
	ID        ID        `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	IsEnabled bool      `json:"is_enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user User) CheckIsEnabled() (enabled bool) {
	return user.IsEnabled
}
