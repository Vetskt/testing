package service

import (
	"context"
	"myapp/config"
	"myapp/graph/model"
)

func CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	db := config.GetDB()

	user := model.User{
		Name:      input.Name,
		Password:  input.Password,
		IsActive:  input.IsActive,
		CreatedAt: *input.CreatedAt,
	}

	err := db.Table("users").Create(&user).Error
	if err != nil {
		panic(err)
	}

	return &user, nil
}

func GetUser(ctx context.Context, id string) (*model.User, error) {
	db := config.GetDB()

	user := model.User{}

	err := db.Table("users").Find(&user, "id = ?", &id).Error
	if err != nil {
		panic(err)
	}

	return &user, nil
}
