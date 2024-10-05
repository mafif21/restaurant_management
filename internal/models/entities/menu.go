package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Menu struct {
	ID          string      `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string      `gorm:"unique;column:name;not null" validate:"required,min=3"`
	Description string      `gorm:"column:description;type:varchar(255)" validate:"max=255"`
	Stock       int         `gorm:"column:stock;not null" validate:"gte=0"`
	IsAvailable bool        `gorm:"column:is_available;not null"`
	Category    []*Category `gorm:"many2many:menu_categories;"`
	Price       float64     `gorm:"column:price;not null" validate:"gt=0"`
	CreatedAt   time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;autoUpdateTime"`
}

func (m *Menu) TableName() string {
	return "menus"
}

func (m *Menu) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.NewString()
	return nil
}
