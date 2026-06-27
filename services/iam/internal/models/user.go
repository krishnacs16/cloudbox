package models


import "time"

type User struct {
	ID        int64
	Username  string
	CreatedAt time.Time
}

type AccessKey struct {
	ID         int64
	UserID     int64
	AccessKey  string
	SecretHash string
	Status     string
	CreatedAt  time.Time
}