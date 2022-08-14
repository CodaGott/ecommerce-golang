package database

import "errors"

var (
	ErrCantFindProduct = errors.New("can't find the product you're looking for")
	ErrCantDecodeProducts = errors.New("can't find product")
	ErrInvalidUserId = errors.New("this user is not valid")
	ErrCantUpdateUser = errors.New("can't add this product to the cart")
	ErrCantRemoveItemCart = errors.New("can't remove this item from the cart")
	ErrCantGetItem = errors.New("was unable to get item fro the cart")
	ErrCantBuyCartItem = errors.New("cannot update the purchase")
)


func AddProductToCart(){
	
}

func RemoveCartItem() {
	
}

func BuyItemFromCart() {
	
}
func InstantBuyer() {

}