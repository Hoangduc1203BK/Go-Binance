package auth

import (
	"binance/api/users"
	constance "binance/const"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func ValidateToken(c *gin.Context) {
	header := strings.Split(c.GetHeader("Authorization"), " ")
	if len(header) < 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
	} else {
		decoded, _ := jwt.Parse(header[1], func(t_ *jwt.Token) (interface{}, error) {
			_, ok := t_.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
			}

			return []byte(TOKEN_KEY), nil
		})

		claims, _ := decoded.Claims.(jwt.MapClaims)
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token expires",
			})
		} else {
			sub := claims["sub"].(float64)
			parseId := uint(sub)
			result, _ := users.ServiceGetUserByID(c, parseId)
			c.Set("user", result)
		}
	}
}
