package requests

type ProductRequest struct {
	ID          string
	Title       string `json:"title" validate:"required,max=190,uniqueProductCategoryTitle" example:"test" `
	Description string `json:"description" validate:"required,max=500" example:"test"`
	SKU         string `json:"sku" validate:"required,max=190" example:"test"`
	Status      string `json:"status" validate:"required" example:"active"`
}

type ProductCategoryRequest struct {
	ID          string
	Title       string `json:"title" validate:"required,max=190,uniqueProductCategoryTitle" example:"test" `
	Description string `json:"description" validate:"required,max=500" example:"test"`
	CoverURL    string `json:"cover_url" validate:"max=190" example:"test"`
	Status      string `json:"status" validate:"required" example:"active"`
}

type ProductBrandRequest struct {
	ID          string
	Title       string `json:"title" validate:"required,max=190,uniqueProductCategoryTitle" example:"test" `
	Description string `json:"description" validate:"required,max=500" example:"test"`
	CoverURL    string `json:"cover_url" validate:"max=190" example:"test"`
	Status      string `json:"status" validate:"required" example:"active"`
}
