package auth

import (
	constance "binance/const"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var TOKEN_KEY = []byte(os.Getenv("SECRET"))

func GenerateToken(userId *uint, typeToken int) (string, error) {
	var exp time.Duration
	if typeToken == constance.TOKEN["AccessToken"] {
		exp = constance.JWT_TIME
	} else {
		exp = constance.REFRESH_TIME
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(exp).Unix(),
	})

	result, err := token.SignedString([]byte(TOKEN_KEY))

	return result, err
}

func DecodeToken(token *string) jwt.MapClaims {
	decoded, _ := jwt.Parse(*token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return TOKEN_KEY, nil
	})

	claims := decoded.Claims.(jwt.MapClaims)

	return claims
}

// func ValidateToken(c *gin.Context) {
// 	header := strings.Split(c.GetHeader("Authorization"), " ")
// 	if header[1] == "undefined" {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 			"error": "Unauthorized",
// 		})
// 	} else {

// 		claims := DecodeToken(&header[1])
// 		if float64(time.Now().Unix()) > claims["exp"].(float64) {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error": "Token expires",
// 			})
// 		} else {
// 			userId, _ := claims["sub"].(float64)
// 			parseId := uint(userId)
// 			result, _ := users.GetUserService(&parseId)
// 			c.Set("user", result)
// 			c.Next()
// 		}

// 	}
// }
