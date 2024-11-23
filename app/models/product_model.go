package models

import "time"

type Product struct {
	BaseModel
	Title             string    `json:"name" gorm:"type:varchar(255);not null"`
	Description       string    `json:"description" gorm:"type:text"`
	DiscountStartDate time.Time `json:"discount_start_date" gorm:"type:timestamp"`
	DiscountEndDate   time.Time `json:"discount_end_date" gorm:"type:timestamp"`
	SellerId          int64     `json:"seller_id" gorm:"not null"`
	Price             float64   `json:"price" gorm:"type:decimal(10,2);not null;default:0;index"`
	PriceDicount      float64   `json:"price_discount" gorm:"type:decimal(10,2);default:0;index"`
	Type              string    `json:"type" gorm:"type:varchar(50);default:'regular';index"`
	Status            string    `json:"status" gorm:"type:varchar(50);default:'draft';index"`
}

type ProductCategory struct {
	BaseNoDateJSONModel
	Title           string `json:"title" gorm:"type:varchar(255);not null;index"`
	Description     string `json:"description" gorm:"type:varchar(500);"`
	CoverUrl        string `json:"cover_url" gorm:"type:varchar(255);"`
	Slug            string `json:"slug" gorm:"type:varchar(255);not null;index"`
	MetaTitle       string `json:"meta_title" gorm:"type:varchar(255);not null"`
	MetaDescription string `json:"meta_description" gorm:"type:varchar(255);not null"`
	MetaKeywords    string `json:"meta_keywords" gorm:"type:varchar(255);not null"`
	Locale          string `json:"locale" gorm:"type:varchar(5);not null"`
}

type ProductBrand struct {
	BaseNoDateJSONModel
	Title           string `json:"title" gorm:"type:varchar(255);not null;index"`
	Description     string `json:"description" gorm:"type:varchar(500);"`
	CoverUrl        string `json:"cover_url" gorm:"type:varchar(255);"`
	Slug            string `json:"slug" gorm:"type:varchar(255);not null;index"`
	MetaTitle       string `json:"meta_title" gorm:"type:varchar(255);not null"`
	MetaDescription string `json:"meta_description" gorm:"type:varchar(255);not null"`
	MetaKeywords    string `json:"meta_keywords" gorm:"type:varchar(255);not null"`
	Locale          string `json:"locale" gorm:"type:varchar(5);not null"`
}
