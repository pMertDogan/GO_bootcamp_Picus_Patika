package jwtUtils

import (
	"encoding/json"
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

//Our token struct
// Reserved claims: https://tools.ietf.org/html/rfc7519#section-4.1
// iss (issuer), *exp (expiration time), sub (subject), aud (audience)
type DecodedJWTToken struct {
	//jsonString: {\"email\":\"pme763@pm.me\",\"exp\":1649622705,\"iat\":1649622105,\"isAdmin\":false,\"isItAccesToken\":true,\"userId\":2}"}
	//\"userId\":2 is decoded as int so we need use it as int not string
	//its selected on login jwt creation step dbUser.ID is int
	UserId         int `json:"userId,required"`
	Email          string `json:"email,required"`
	Iat            int    `json:"iat,required"` //issued at  *optional
	Exp            int64  `json:"exp,required"` //expiration time *Must be used
	IsAdmin        bool   `json:"isAdmin,required"`
	IsItAccesToken bool   `json:"isItAccesToken,required"`
	// Iss    string   `json:"iss"`
}

/*
https://auth0.com/learn/json-web-tokens/
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)

  header.payload.signature

*/
func GenerateToken(claims *jwt.Token, secret string) string {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ := claims.SignedString(hmacSecret)

	return token
}

/*
This Method is verify token checks exp dates too! Then bind token to struct
https://pkg.go.dev/github.com/golang-jwt/jwt/v4@v4.4.1?utm_source=gopls#Parse

*/
func VerifyDecodeToken(token string, secret string) (*DecodedJWTToken, error) {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("Token is malformed")
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			// Token is either expired or not active yet
			return nil, errors.New("Token is either expired or not active yet")
		} else {
			return nil, errors.New("Couldn't handle this token: " + err.Error())
		}
	}

	var decodedToken DecodedJWTToken
	zap.L().Sugar().Debug("Decoded token: ", decoded)
	decodedClaims := decoded.Claims.(jwt.MapClaims)
	jsonString, _ := json.Marshal(decodedClaims)
	zap.L().Sugar().Debug("jsonString: ", string(jsonString))
	json.Unmarshal(jsonString, &decodedToken)
	zap.L().Sugar().Debug("decodedToken: ", decodedToken)
	return &decodedToken, nil
}
