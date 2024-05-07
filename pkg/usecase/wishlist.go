package usecase

import (
	"CrocsClub/pkg/repository/interfaces"
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"
)

type WishlistUsecase struct {
	repo interfaces.WishlistRepository
}

func NewWishlistUsecase(repo interfaces.WishlistRepository) services.WishlistUsecase {
	return &WishlistUsecase{
		repo: repo,
	}
}

func (w *WishlistUsecase) AddToWishlist(userId, invId int) error {
	ok, err := w.repo.CheckIfProductExist(invId, userId)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("product already exist")
	}
	err = w.repo.AddToWishlist(userId, invId)
	if err != nil {
		return errors.New("cannot add to database")
	}
	return nil
}
func (w *WishlistUsecase) GetWishlist(userId int) ([]models.Wishlist, error) {
	data, err := w.repo.GetWishlist(userId)
	if err != nil {
		return []models.Wishlist{}, errors.New("error in fetching details")
	}
	return data, nil
}
func (w *WishlistUsecase) RemoveFromWishlist(userId, invId int) error {
	err := w.repo.RemoveFromWishlist(userId, invId)
	if err != nil {
		return errors.New("error in removing data")
	}
	return nil
}
