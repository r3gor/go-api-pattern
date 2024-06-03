package service

import "fmt"

type BuyService struct{}

func (s BuyService) Buy(userId string, cartId string, amount int32) {
	fmt.Println("Cart #", cartId, "| Buyed by:", userId)
}
