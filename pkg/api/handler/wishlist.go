package handler

import (
	"CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	usecase interfaces.WishlistUsecase
}

func NewWishlistHandler(usecase interfaces.WishlistUsecase) *WishlistHandler {
	return &WishlistHandler{
		usecase: usecase,
	}
}
func (w *WishlistHandler) AddToWishlist(cxt *gin.Context) {
	userId, ok := cxt.Get("id")
	uId := userId.(int)
	if !ok {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot add to wishlist", nil, errors.New("error getting id"))
		cxt.JSON(http.StatusBadRequest, errRes)
		return
	}
	inventory_Id := cxt.Query("inventory_id")
	invId, err := strconv.Atoi(inventory_Id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		cxt.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = w.usecase.AddToWishlist(uId, invId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot add to wishlist", nil, err.Error())
		cxt.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully added wishlist", nil, nil)
	cxt.JSON(http.StatusOK, successRes)

}
func (w *WishlistHandler) GetWishlist(cxt *gin.Context) {
	userId, ok := cxt.Get("id")
	uId := userId.(int)
	if !ok {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot get wishlist", nil, errors.New("error getting id"))
		cxt.JSON(http.StatusBadRequest, errRes)
		return
	}
	wishlist, err := w.usecase.GetWishlist(uId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot get wishlist", nil, err.Error())
		cxt.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got wishlist", wishlist, nil)
	cxt.JSON(http.StatusOK, successRes)
}

func (w *WishlistHandler) RemoveFromWishlist(cxt *gin.Context) {
	userId, ok := cxt.Get("id")
	uId := userId.(int)
	if !ok {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot remove from wishlist", nil, errors.New("error getting id"))
		cxt.JSON(http.StatusBadRequest, errRes)
		return
	}
	inventory_Id := cxt.Query("inventory_id")
	invId, err := strconv.Atoi(inventory_Id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "conversion failed", nil, err.Error())
		cxt.JSON(http.StatusBadRequest, errRes)
		return
	}
	err = w.usecase.RemoveFromWishlist(uId, invId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot remove from wishlist", nil, err.Error())
		cxt.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully removed from wishlist", nil, nil)
	cxt.JSON(http.StatusOK, successRes)
}
