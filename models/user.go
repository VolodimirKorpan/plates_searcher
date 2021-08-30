package models

import "time"

// User is a JSON user
type User struct {
	ID        int      `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (user *User) ToDB() *DBUser {
	return &DBUser{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}

type DBUser struct {
	ID        int `gorm:"primaryKey"`
	Username  string
	Password  string
	CreatedAt time.Time
}


// ToWeb converts DBUser to User
func (dbUser *DBUser) ToWeb() *User {
	return &User{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Password:  dbUser.Password,
		CreatedAt: dbUser.CreatedAt,
	}
}
