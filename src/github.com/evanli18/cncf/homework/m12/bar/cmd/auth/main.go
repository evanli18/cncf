package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"bar/configs"
	"bar/pkg/auth"
	"bar/pkg/server"
	"bar/pkg/user"

	"github.com/dgrijalva/jwt-go"
)

var JWTPrivateKey = `-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBANjCxru4FLhVt0YO
m+zjMuQ1Vtnhz/3C+H052HeHnE0bMw3v4NlGGekR+gBTQg7fC/DWUyz9oGTuvqte
0nh9ixotfpR99chWlGFRgPKW5FObK0DrDRlAQwmp6ROwMDhFCLyUJazajuNL09Fq
JSkilatKrB3A0FmRlWPeCBqy1UHVAgMBAAECgYB/IPV3wYX9euBLqWPP8oy1hYcT
sLnRBhnBMD0CFboZCvvNj8PbCp9Fr/JlYG9c03poXPtZZsM8jz2quqlMW61Jr0GH
XfmfLxBxKDZKd4m9uwz1f0oiR5gbHL6geKanp9lOaEEQ5UhfPeLIh4eyVmf5FmKe
pFD6aEL2qMaqbsaGgQJBAPtNtLUE5Dfeeg835+sNgvvqFB9NHPNhtBlg1FHtVxr6
BuYtfsv2/FpZgWXqYukwNlxUHJIZw4VoX6j6qz3AeaECQQDcz89pIiSAAYdtSZ3Z
vzBxyYGy+DiM/Txo6YF2cZsWpKFUKol8kMx+Y3VfLH0shR3Fyp0vmEwSEVSO51t6
9mO1AkB5UCC9Bfh5s+9ua0mMocAqhexi0+H256J+Ycz9I7rZ7frooOvF4JwfrXeW
0FghQ8HqPjxwlvlY7HLJawDBVaohAkA/1K7vhFgqzMZaWFqSNIuLiSW+F7U5RIcv
CLlNBQBBJmwgiX9fC/ihXJz0W0cAFKcLo0uXE56B5pKcENNIE2u1AkEApE9py2q4
DERIc6uBEa+PVA4SCktQ3OGW4IdUe6V1wXv/1qxQNYFuChyNL4vO/m2c01LysJdf
8HdviLjAu0N53w==
-----END PRIVATE KEY-----`

func createJwtToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat":      time.Now().Unix(),                     // Token颁发时间
		"nbf":      time.Now().Unix(),                     // Token生效时间
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token过期时间，目前是24小时
		"iss":      "example.com",                         // 颁发者
		"sub":      "AccessKey",                           // 主题
		"username": username,
	})
	key, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(JWTPrivateKey))
	return token.SignedString(key)
}

func httpHandleAuth(request *http.Request) server.Response {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return server.Response{Code: 500, Body: err.Error()}
	}

	var user user.User
	if err := json.Unmarshal(body, &user); err != nil {
		return server.Response{Code: 500, Body: err.Error()}
	}

	send, _ := json.Marshal(user)
	resp, err := http.Post(configs.USERSERVER+"/login", "application/json", bytes.NewBuffer(send))
	if err != nil {
		log.Println(err.Error())
		return server.Response{Code: 500, Body: "call user service err."}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		token, err := createJwtToken(user.Username)
		if err != nil {
			return server.Response{Code: 200, Body: err.Error()}
		}
		return server.Response{Code: 401, Body: token}
	}

	return server.Response{Code: 401, Body: "auth error"}
}

func httpHandleVerify(request *http.Request) server.Response {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return server.Response{Code: 500, Body: err.Error()}
	}

	username, err := auth.Verify(string(body))
	if err != nil {
		return server.Response{Code: 401, Body: err.Error()}
	}
	return server.Response{Code: 200, Body: username}
}

func main() {
	ser := server.New()
	ser.Route("/auth", httpHandleAuth)
	ser.Route("/verify", httpHandleVerify)

	log.Fatal(ser.Run(":8082"))
}
