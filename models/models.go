package models

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
	"time"
)

type Pokemon struct {
	ID        uint `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string `json:"name"`
	PkdxId    uint `json:"pkdxId"`
	Types     pgtype.JSONB `json:"types"`
}
