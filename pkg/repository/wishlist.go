package repository

import (
	"CrocsClub/pkg/repository/interfaces"
	"CrocsClub/pkg/utils/models"

	"gorm.io/gorm"
)

type wishlistRepository struct {
	DB *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) interfaces.WishlistRepository {
	return &wishlistRepository{
		DB: db,
	}
}
func (aw *wishlistRepository) AddToWishlist(userId, invId int) error {
	query := "insert into wishlists(inventory_id, user_id) values (?,?)"
	err := aw.DB.Exec(query, invId, userId).Error
	if err != nil {
		return err
	}
	return nil
}
func (aw *wishlistRepository) CheckIfProductExist(invId, userId int) (bool, error) {
	var count int
	query := "select count(*) from wishlists where inventory_id = ? and user_id = ?"
	err := aw.DB.Raw(query, invId, userId).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (aw *wishlistRepository) GetWishlist(userId int) ([]models.Wishlist, error) {
	query := "select inventory_id from wishlists where user_id = ?"
	var wishlist []models.Wishlist
	err := aw.DB.Raw(query, userId).Scan(&wishlist).Error
	if err != nil {
		return []models.Wishlist{}, err
	}
	return wishlist, nil
}
func (aw *wishlistRepository) RemoveFromWishlist(userId, invId int) error {
	query := "delete from wishlists where user_id = ? and inventory_id = ?"
	err := aw.DB.Exec(query, userId, invId).Error
	if err != nil {
		return err
	}
	return nil

}
