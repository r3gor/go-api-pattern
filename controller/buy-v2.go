package controller

import (
	"api-pattern-go/service"
	"context"
)

type BuyController struct {
	BuySvc service.BuyService
}

type BuyPayload struct {
	CartId *string `json:"cartId"`
	Amount *int32  `json:"amount"`
}

func (c BuyController) HandleBuy(ctx context.Context, payload interface{}) {
	p, ok := payload.(*BuyPayload)
	if !ok {
		panic("payload is not BuyPayload")
	}
	c.BuySvc.Buy("roger", *p.CartId, *p.Amount)
}

func (c BuyController) HandleBuyWithPromo(ctx context.Context, payload interface{}) {
	p, ok := payload.(*BuyPayload)
	if !ok {
		panic("payload is not BuyPayload")
	}
	c.BuySvc.Buy("roger", *p.CartId, *p.Amount)
}

type ControllerFunc func(ctx context.Context, payload interface{})

func (f ControllerFunc) Handle(ctx context.Context, payload interface{}) {
	f(ctx, payload)
}
