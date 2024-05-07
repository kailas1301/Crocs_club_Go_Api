package interfaces

import "CrocsClub/pkg/utils/models"

type WishlistUsecase interface {
	AddToWishlist(userId, invId int) error
	GetWishlist(userId int) ([]models.Wishlist, error)
	RemoveFromWishlist(userId, invId int) error
}
