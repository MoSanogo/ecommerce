package repository

import (
	"ecommerce-grpc-api/internal/dtos"
	"ecommerce-grpc-api/internal/models"
)

type UserRepository interface {
	InsertOne(data dtos.UserCreateDto) string
	GetOne(id string) *models.User
	UpdateOne(id string, data dtos.UserUpdateDto) bool
	DeleteOne(id string) bool
	GetAll(searchCriteria dtos.SearchCriteria) []*models.User
}
