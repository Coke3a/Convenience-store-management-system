package repository_promotion

import (
	"gorm.io/gorm"
	"time"
)

type promotionRepository struct {
	db *gorm.DB
}

type Promotion struct {
	ID             string `gorm:"primary_key;not_null"`
	Name           string
	Type           int
	PurchaseMin    int
	Required_point int
	Discount       int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewPromotionRepository(db *gorm.DB) PromotionRepository {
	db.Table("promotions").AutoMigrate(&Promotion{})
	return promotionRepository{db: db}
}

type PromotionRepository interface {
	Create(promotion Promotion) error
	Update(promotion Promotion) error
	Delete(id string) error
	FindAll() (promotions []Promotion, err error)
	FindbyId(id string) (promotion Promotion, err error)
}
