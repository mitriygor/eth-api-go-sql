package models

import "time"

type Timestamps struct {
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime"`
}
