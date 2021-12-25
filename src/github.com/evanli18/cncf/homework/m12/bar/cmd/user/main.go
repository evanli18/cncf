package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"bar/pkg/auth"
	"bar/pkg/server"
	"bar/pkg/user"
)

// make test users
type users struct {
	u    []user.User
	lock sync.Mutex
}

var mockUsers = users{
	u: []user.User{
		{
			Username: "admin",
			Password: "admin",

			SuperAdmin: true,
		},
		{
			Username: "evan",
			Password: "evan",

			SuperAdmin: false,
		},
	},
	lock: sync.Mutex{},
}

func httpHandleLogin(request *http.Request) server.Response {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return server.Response{Code: 500, Body: err.Error()}
	}

	var user user.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return server.Response{Code: 500, Body: err.Error()}
	}

	mockUsers.lock.Lock()
	defer mockUsers.lock.Unlock()

	for _, loopU := range mockUsers.u {
		if user.Username == loopU.Username && user.Password == loopU.Password {
			return server.Response{Code: 200, Body: "ok"}
		}
	}

	return server.Response{Code: 404, Body: "not found"}
}

func httpHandleProfile(http *http.Request) server.Response {
	username, err := auth.VerifyFromHeader(http.Header)
	if err != nil {
		return server.Response{Code: 401, Body: err.Error()}
	}
	for _, u := range mockUsers.u {
		if u.Username == username {
			body, _ := json.Marshal(u)
			return server.Response{Code: 200, Body: string(body)}
		}
	}
	return server.Response{Code: 404, Body: "user not found"}
}

func main() {
	ser := server.New()
	ser.Route("/login", httpHandleLogin)
	ser.Route("/profile", httpHandleProfile)

	log.Fatal(ser.Run(":8081"))
}
