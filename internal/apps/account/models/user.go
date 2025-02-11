package models

import "time"

type HashOptions struct {
	Type       string `gorm:"type:varchar(255)" json:"type"`
	MemoryCost int    `gorm:"type:int" json:"memoryCost"`
	TimeCost   int    `gorm:"type:int" json:"timeCost"`
	Threads    int    `gorm:"type:int" json:"threads"`
}

type Target struct {
	ID           string    `gorm:"primaryKey;column:$id;type:uuid;default:gen_random_uuid()" json:"$id"`
	CreatedAt    time.Time `gorm:"column:$createdAt" json:"$createdAt"`
	UpdatedAt    time.Time `gorm:"column:$updatedAt" json:"$updatedAt"`
	Name         string    `gorm:"type:varchar(255)" json:"name"`
	UserID       string    `gorm:"type:uuid" json:"userId"`
	ProviderID   string    `gorm:"type:varchar(255)" json:"providerId"`
	ProviderType string    `gorm:"type:varchar(255)" json:"providerType"`
	Identifier   string    `gorm:"type:varchar(255)" json:"identifier"`
	Expired      bool      `gorm:"type:boolean" json:"expired"`
}

type User struct {
	ID                string      `gorm:"primaryKey;column:$id;type:uuid;default:gen_random_uuid()" json:"$id"`
	CreatedAt         time.Time   `gorm:"column:$createdAt" json:"$createdAt"`
	UpdatedAt         time.Time   `gorm:"column:$updatedAt" json:"$updatedAt"`
	Name              string      `gorm:"type:varchar(255)" json:"name"`
	Password          string      `gorm:"type:text" json:"password"`
	Hash              string      `gorm:"type:varchar(255)" json:"hash"`
	HashOptions       HashOptions `gorm:"embedded" json:"hashOptions"`
	Registration      time.Time   `gorm:"type:timestamp" json:"registration"`
	Status            bool        `gorm:"type:boolean" json:"status"`
	Labels            []string    `gorm:"type:text[]" json:"labels"`
	PasswordUpdate    time.Time   `gorm:"type:timestamp" json:"passwordUpdate"`
	Email             string      `gorm:"type:varchar(255)" json:"email"`
	Phone             string      `gorm:"type:varchar(255)" json:"phone"`
	EmailVerification bool        `gorm:"type:boolean" json:"emailVerification"`
	PhoneVerification bool        `gorm:"type:boolean" json:"phoneVerification"`
	MFA               bool        `gorm:"type:boolean" json:"mfa"`
	Prefs             interface{} `gorm:"type:jsonb" json:"prefs"`
	Targets           []Target    `gorm:"foreignKey:UserID" json:"targets"`
	AccessedAt        time.Time   `gorm:"type:timestamp" json:"accessedAt"`
}
