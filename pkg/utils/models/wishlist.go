package models

type Wishlist struct {
	UserID      int `json:"user_id"`
	InventoryID int `json:"inventory_id"`
}
