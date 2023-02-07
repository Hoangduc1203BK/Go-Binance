package auth

import (
	"fmt"
	"binance/database"
	"binance/model"
)

func CollectionInsertToken(data *model.Token) model.Token {
	token := model.Token{TokenString: data.TokenString, UserId: data.UserId, Blacklisted: data.Blacklisted, Expires: data.Expires}
	database.DB.Create(&token)

	fmt.Println(token)
	return token
}

func CollectionFindToken(token *string) model.Token {
	var result model.Token
	database.DB.Where("token_string = ?", token).First(&result)

	return result
}

func CollectionFindAllToken(userId *uint, blackList *bool) []model.Token {
	var tokens []model.Token
	database.DB.Where("user_id = ? AND blacklisted = ?", *userId, *blackList).Find(&tokens)
	fmt.Println(tokens)

	return tokens
}

func CollectionUpdateToken(id *uint, blacklisted *bool) {
	var token model.Token

	database.DB.Model(&token).Where("id =  ?", id).Updates(model.Token{
		Blacklisted: true,
	})
}

func CollectionDeleteToken(token *string) error {
	var tokenString model.Token
	result := database.DB.Where("token_string = ?", *token).Delete(&tokenString)

	return result.Error
}
