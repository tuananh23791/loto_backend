package model

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	UserId string `json:user_id`
	Rold   string `json:role`
	jwt.StandardClaims
}
