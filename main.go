package main

import (
	"api-pattern-go/controller"
	"api-pattern-go/middleware"
	"api-pattern-go/port"
	"api-pattern-go/service"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

type AppHandler struct {
	Controller      port.Controller
	MiddlewareChain []port.Middleware
	Pipe            interface{}
	Payload         interface{}
}

func (h *AppHandler) BindPayload(r *http.Request) {
	t := reflect.TypeOf(h.Pipe)
	h.Payload = reflect.New(t).Interface()
	json.NewDecoder(r.Body).Decode(h.Payload)
}

func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	h.BindPayload(r)
	fmt.Println("AppHandler.Payload:", h.Payload)

	for _, mw := range h.MiddlewareChain {
		mw.Handle(ctx, w, r)
	}

	h.Controller.Handle(ctx, h.Payload)
}

func main() {
	authService := service.AuthService{}
	buyService := service.BuyService{}
	authMiddleware := middleware.CheckSession{
		AuthSvc: authService,
	}
	buyController := controller.BuyController{
		BuySvc: buyService,
	}

	http.Handle("/foo", AppHandler{
		Pipe:            controller.BuyPayload{},
		MiddlewareChain: []port.Middleware{authMiddleware},
		Controller:      controller.ControllerFunc(buyController.HandleBuy),
	})

	port := fmt.Sprintf(":%d", 3000)
	log.Println("Server running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
