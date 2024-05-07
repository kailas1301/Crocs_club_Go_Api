package domain

type Wishlist struct {
	ID          uint  `json:"id" gorm:"primarykey"`
	UserID      uint  `json:"user_id"`
	User        Users `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	InventoryID int   `json:"inventory_id" gorm:"not null"`
}
