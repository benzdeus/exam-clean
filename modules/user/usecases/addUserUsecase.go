package usecases

import "demo-clean/entitiies"

func (u *userUsecase) AddUser(user *entitiies.UserModel) (*entitiies.UserModel, error) {
	return u.repo.InsertUser(user)
}
