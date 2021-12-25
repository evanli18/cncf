package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"bar/configs"
	"bar/pkg/server"
	"bar/pkg/user"
)

type Order struct {
	ID    int64
	Goods []string
}

type MyOrders struct {
	Username   string
	Superadmin bool

	Oders []Order
}

// httpHandleMyOrders 组装自己的订单数据
// from user, order
func httpHandleMyOrders(request *http.Request) server.Response {
	var user user.User
	var orders []Order

	// new header
	var header = http.Header{}
	for k, v := range request.Header {
		k = strings.ToLower(k)
		if strings.HasPrefix(k, "x-") {
			header.Set(k, v[0])
		}
	}
	header.Set("Authorization", request.Header.Get("Authorization")) //auth

	// 查用户信息
	req, _ := http.NewRequest("GET", configs.USERSERVER+"/profile", nil)
	req.Header = header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return server.Response{Code: 500, Body: "call user service err"}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != 200 {
		return server.Response{Code: resp.StatusCode, Body: "user err"}
	}
	if err := json.Unmarshal(body, &user); err != nil {
		return server.Response{Code: 500, Body: err.Error()}
	}

	// 查订单信息
	req, _ = http.NewRequest("GET", configs.ORDERSERVER+"/myorder", nil)
	req.Header = header
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return server.Response{Code: 500, Body: "call order service err"}
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode != 200 {
		return server.Response{Code: resp.StatusCode, Body: err.Error()}
	}
	if err := json.Unmarshal(body, &orders); err != nil {
		return server.Response{Code: 500, Body: err.Error()}
	}

	data, _ := json.Marshal(MyOrders{
		Username:   user.Username,
		Superadmin: user.SuperAdmin,
		Oders:      orders,
	})
	return server.Response{Code: 200, Body: string(data)}
}

func main() {
	ser := server.New()
	ser.Route("/myorders", httpHandleMyOrders)

	log.Fatal(ser.Run(":8084"))
}
