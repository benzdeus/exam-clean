package repositories

import "demo-clean/entitiies"

func (r *userRepo) InsertUser(user *entitiies.UserModel) (*entitiies.UserModel, error) {
	result := r.db.Create(&user)

	return user, result.Error

}
