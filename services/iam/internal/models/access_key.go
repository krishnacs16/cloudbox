package models

type AccessKeyStatus string

const (
	AccessKeyActive   AccessKeyStatus = "ACTIVE"
	AccessKeyInactive AccessKeyStatus = "INACTIVE"
)


type AccessKey struct {
	BaseModel

	UserID int64

	AccessKey string

	SecretHash string

	Status AccessKeyStatus
}