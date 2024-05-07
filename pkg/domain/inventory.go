package domain

type Category struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Category string `json:"category"`
}

type Inventories struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
	ProductName string   `json:"product_name"`
	Size        string   `json:"size" gorm:"size:W4;default:'W6';Check:size IN ('W4', 'W5', 'W6', 'W7', 'W8','W9', 'W10', 'M4', 'M5', 'M6', 'M7', 'M8', 'M9', 'M10', 'K1', 'K2', 'K3', 'K4');"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
	// Image       string   `json:"product_image"`
}

type Image struct {
	ID          int         `json:"id"`
	URL         string      `json:"url"`
	InventoryID int         `json:"inventory_id"`
	Inventories Inventories `json:"_" gorm:"foreignKey:InventoryID"`
}
