package entitiies

type UserRepo interface {
	InsertUser(*UserModel) (*UserModel, error)
}

type UserUsecase interface {
	AddUser(*UserModel) (*UserModel, error)
}

type UserModel struct {
	ID       int64  `json:"id" gorm:"increment"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (UserModel) TableName() string {
	return "users"
}
