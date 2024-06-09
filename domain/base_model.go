package domain

import (
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type BaseModel struct {
	Union_ID  string    `gorm:"type:varchar(8);column:union_id;not null" json:"union_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// BeforeCreate will set a UUID. This method is a Gorm hook triggered before creating a new record.
func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if base.Union_ID == "" {
		id, err := nanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-", 8)
		if err != nil {
			return err
		}
		base.Union_ID = id
	}
	return nil
}
