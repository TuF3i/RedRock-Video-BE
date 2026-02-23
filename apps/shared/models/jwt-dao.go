package models

import "github.com/golang-jwt/jwt/v5"

type MainClaims struct {
	Uid  int64  `json:"uid"`
	Role string `json:"role"`
	Type string `json:"type"`

	jwt.RegisteredClaims
}
