package port

import (
	"context"
	"net/http"
)

type CartService interface {
	BuyItem(id string) error
}

type AuthService interface {
	CheckSession(token string) (bool, error)
	Login(username string, password string) (string, error)
}

type Middleware interface {
	Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) error // TODO: quitar w de los parametros
}

type Controller interface {
	Handle(ctx context.Context, payload interface{})
}
