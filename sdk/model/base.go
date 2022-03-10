package model

import (
	"time"
)

type BaseModel struct {
	ID        int64     `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
