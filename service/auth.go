package service

import "fmt"

type AuthService struct{}

func (s AuthService) CheckSession(token string) (bool, error) {
	fmt.Println("CheckSession")
	return true, nil
}

func (s AuthService) Login(username string, password string) (string, error) {
	return "token1234", nil
}
