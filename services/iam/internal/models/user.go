package models

type User struct {
	BaseModel

	Username string
}

type AccessKey struct {
	ID         int64
	UserID     int64
	AccessKey  string
	SecretHash string
	Status     string
	CreatedAt  time.Time
}