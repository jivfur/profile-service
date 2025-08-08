package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID                string     `gorm:"type:char(36);primaryKey"`      // UUID format
	UID               *string    `gorm:"type:varchar(255);uniqueIndex"` // Firebase UID, nullable if used
	Username          string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	Email             string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	PasswordHash      string     `gorm:"type:varchar(255);not null"` // Store hashed password
	Name              string     `gorm:"type:varchar(255);not null"`
	Bio               *string    `gorm:"type:text"`
	Gender            string     `gorm:"type:enum('MALE','FEMALE','TRANS MAN','TRANS WOMAN','NON BINARY','QUEER');not null"`
	SexualOrientation string     `gorm:"type:enum('GAY','STRAIGHT','BISEXUAL','ASEXUAL','DEMISEXUAL','PANSEXUAL','QUEER');not null"`
	SexualPosition    string     `gorm:"type:enum('TOP','VERS TOP','VERSATILE','VERS BOTTOM','BOTTOM','SIDE');not null"`
	PhotoURL          *string    `gorm:"type:varchar(1024)"`
	Location          *string    `gorm:"type:varchar(255)"` // Optional text location
	DateOfBirth       *time.Time `gorm:"type:date"`         // Use date type for DOB
	Hobbies           *string    `gorm:"type:text"`         // Comma-separated or JSON
	Interests         *string    `gorm:"type:text"`
	EmailVerified     bool       `gorm:"default:false"`
	FaceVerified      bool       `gorm:"default:false"`
	VerificationScore float32    `gorm:"default:0"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"` // Soft delete
	Photos            []Photo        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Locations         []Location     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Photo struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UserID       string    `gorm:"type:char(36);index;not null"`
	URL          string    `gorm:"type:varchar(1024);not null"`
	Caption      *string   `gorm:"type:text"`
	IsPublic     bool      `gorm:"default:true"`
	IsNSFW       bool      `gorm:"default:false"`
	PhotoOrder   int       `gorm:"default:0"`
	DateUploaded time.Time `gorm:"autoCreateTime"`
}

type Location struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    string    `gorm:"type:char(36);index;not null"`
	Latitude  float64   `gorm:"not null"`
	Longitude float64   `gorm:"not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

var Genders = map[string]bool{
	"MALE":        true,
	"FEMALE":      true,
	"TRANS MAN":   true,
	"TRANS WOMAN": true,
	"NON BINARY":  true,
	"QUEER":       true,
}

var SexualOrientation = map[string]bool{
	"GAY":        true,
	"STRAIGHT":   true,
	"BISEXUAL":   true,
	"ASEXUAL":    true,
	"DEMISEXUAL": true,
	"PANSEXUAL":  true,
	"QUEER":      true,
}

var SexualPosition = map[string]bool{
	"TOP":         true,
	"VERS TOP":    true,
	"VERSATILE":   true,
	"VERS BOTTOM": true,
	"BOTTOM":      true,
	"SIDE":        true,
}
