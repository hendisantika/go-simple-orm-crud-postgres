package models

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	PaymentCode string
	Name        string
	Status      string
	Created     time.Time `gorm:"autoCreateTime"`
	Updated     time.Time `gorm:"autoUpdateTime"`
}
