package usecases

import "demo-clean/entitiies"

type userUsecase struct {
	repo entitiies.UserRepo
}

func NewUserUsecase(repo entitiies.UserRepo) *userUsecase {
	return &userUsecase{
		repo,
	}
}
