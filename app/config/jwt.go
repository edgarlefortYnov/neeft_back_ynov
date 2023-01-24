package config

/**
 *@Author: ANYARONKE Dare Samuel
 */

import "github.com/golang-jwt/jwt/v4"

var JWT_SECRET = []byte("aqwzsxedcrfvtgbyhnujujikolpmamzlekjhgfdswqazx")

//var COOKIE_TOKEN = "token"

type JWTClaims struct {
	UserId    uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	jwt.RegisteredClaims
}
