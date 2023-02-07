package auth

import (
	"fmt"
	constance "binance/const"
	"binance/model"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(*password), 10)

	return string(hashPassword), err
}

func ComparePassword(password *string, hashPassword *string) error {
	err := bcrypt.CompareHashAndPassword([]byte(*hashPassword), []byte(*password))
	return err
}

// func LogInService(data *users.GetUserDto) (TokenType, error) {
// 	result, err := users.GetUserByAuth(data)

// 	token := TokenType{
// 		accessToken:  "",
// 		refreshToken: "",
// 	}

// 	if err != nil {
// 		return token, err
// 	} else {
// 		// generate access token
// 		access, accessErr := GenerateToken(&result.ID, 1)

// 		if accessErr != nil {
// 			return token, accessErr
// 		}
// 		token.accessToken = access

// 		// generate refresh token
// 		refresh, refreshErr := GenerateToken(&result.ID, 2)
// 		if refreshErr != nil {
// 			return token, refreshErr
// 		}
// 		token.refreshToken = refresh

// 		// save token to DB
// 		exp := time.Now().Add(constance.REFRESH_TIME).Unix()
// 		doc := model.Token{TokenString: refresh, UserId: result.ID, Blacklisted: false, Expires: exp}
// 		tokenResult := CollectionInsertToken(&doc)

// 		if tokenResult.ID == 0 {
// 			return token, fmt.Errorf("Fail to insert refreshToken to db: %v", refresh)
// 		}

// 		return token, nil
// 	}
// }

func RefreshTokenService(refreshToken *string) (TokenType, error) {
	decoded := DecodeToken(refreshToken)
	tokenResponse := TokenType{
		accessToken:  "",
		refreshToken: "",
	}
	if float64(time.Now().Unix()) > decoded["exp"].(float64) {
		return tokenResponse, fmt.Errorf("Token expires")
	} else {
		token := CollectionFindToken(refreshToken)

		if token.Blacklisted == true {
			userId := decoded["sub"].(float64)
			parseId := uint(userId)
			blackList := false
			tokens := CollectionFindAllToken(&parseId, &blackList)

			for _, t := range tokens {
				blacklisted := false
				CollectionUpdateToken(&t.ID, &blacklisted)
			}
			return tokenResponse, fmt.Errorf("Invalid refresh token")
		} else {
			blacklisted := false
			CollectionUpdateToken(&token.ID, &blacklisted)
			userId := decoded["sub"].(float64)
			parseId := uint(userId)

			// generate new access token
			access, accessErr := GenerateToken(&parseId, 1)

			if accessErr != nil {
				return tokenResponse, accessErr
			}
			tokenResponse.accessToken = access

			// generate new refresh token
			refresh, refreshErr := GenerateToken(&parseId, 2)
			if refreshErr != nil {
				return tokenResponse, refreshErr
			}
			tokenResponse.refreshToken = refresh

			// save token to DB
			exp := time.Now().Add(constance.REFRESH_TIME).Unix()
			doc := model.Token{TokenString: refresh, UserId: parseId, Blacklisted: false, Expires: exp}
			tokenResult := CollectionInsertToken(&doc)

			if tokenResult.ID == 0 {
				return tokenResponse, fmt.Errorf("Fail to insert refreshToken to db: %v", refresh)
			}

			return tokenResponse, nil
		}
	}
}

func LogOutService(refreshToken *string) error {
	token := CollectionFindToken(refreshToken)

	if token.Blacklisted == true {
		return fmt.Errorf("Invalid refresh token")
	}

	err := CollectionDeleteToken(refreshToken)
	if err != nil {
		return fmt.Errorf("Fail to delete token: %v", *refreshToken)
	}

	return nil

}
