package main

import (
	"encoding/json"
	"log"
	"net/http"

	"bar/pkg/auth"
	"bar/pkg/server"
)

type Order struct {
	ID    int64
	Goods []string
}

var oders = map[string][]Order{
	"admin": {
		{
			ID:    1,
			Goods: []string{"Beer", "Whisky"},
		},
	},
}

func httpHandleQueryMyOrders(request *http.Request) server.Response {
	outOrders := make([]Order, 0)
	user, err := auth.VerifyFromHeader(request.Header)
	if err != nil {
		return server.Response{Code: 401, Body: err.Error()}
	}

	if _, ok := oders[user]; ok {
		outOrders = append(outOrders, oders[user]...)
	}

	data, _ := json.Marshal(outOrders)
	return server.Response{Code: 200, Body: string(data)}
}

func main() {
	ser := server.New()
	ser.Route("/myorder", httpHandleQueryMyOrders)

	log.Fatal(ser.Run(":8083"))
}
