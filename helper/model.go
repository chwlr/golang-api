package helper

import (
	"github.com/chwlr/golang-api/model/domain"
	"github.com/chwlr/golang-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id: 	category.Id,
		Name: 	category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id: user.Id,
		IdRole: user.IdRole,
		Name: user.Name,
		Email: user.Email,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}