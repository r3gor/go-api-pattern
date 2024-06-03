package controller

// type BuyController struct {
// 	BuySvc service.BuyService
// }
//
// type BuyPayload struct {
// 	CartId *string `json:"cartId"`
// 	Amount *int32  `json:"amount"`
// }
//
// func (c BuyController) Handle(ctx context.Context, payload interface{}) {
// 	p, ok := payload.(*BuyPayload)
// 	if !ok {
// 		panic("payload is not BuyPayload")
// 	}
// 	c.BuySvc.Buy("roger", *p.CartId, *p.Amount)
// }
