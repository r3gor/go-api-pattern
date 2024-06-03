package middleware

import (
	"api-pattern-go/port"
	"context"
	"fmt"
	"net/http"
)

type CheckSession struct {
	AuthSvc port.AuthService
}

func (m CheckSession) Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	token := "Bearer1234"
	validSession, err := m.AuthSvc.CheckSession(token)
	if err != nil {
		return err
	}
	if !validSession {
		ctx.Done()
	}
	fmt.Println("Auth Middleware")
	return nil
}
