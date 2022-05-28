package usersstore

import "time"

type Token struct {
	AccessToken  string    `bson:"access_token,omitempty"`
	Expiry       time.Time `bson:"expiry,omitempty"`
	RefreshToken string    `bson:"refresh_token,omitempty"`
	TokenType    string    `bson:"token_type,omitempty"`
}

func NewToken() *Token {
	return &Token{}
}
