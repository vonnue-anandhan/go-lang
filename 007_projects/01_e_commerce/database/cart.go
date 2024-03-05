package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrCantFindCartItem    = errors.New("can't find the product")
	ErrCantDecodeCartItems = errors.New("can't find the product")
	ErrUserIdIsNotValid    = errors.New("this user is not valid")
	ErrCantUpdateUser      = errors.New("cannot add the product to the cart")
	ErrCantRemoveCartItem  = errors.New("cannot remove this item from the cart")
	ErrCantGetCartItem     = errors.New("unable to get the item from the cart")
	ErrCantCartBuyItem     = errors.New("cannot update the purchase")
)

func AddProductToCart() gin.HandlerFunc {}

func RemoveCartItem() gin.HandlerFunc {}

func BuyItemFromCart() gin.HandlerFunc {}

func InstantBuy() gin.HandlerFunc {}
