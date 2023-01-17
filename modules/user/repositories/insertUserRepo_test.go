package repositories_test

import (
	"demo-clean/db"
	"demo-clean/entitiies"
	"demo-clean/modules/user/repositories"
	"log"
	"testing"
)

func TestInsertRepo(t *testing.T) {
	dbPg, tearDown := db.New()

	defer tearDown()

	userRepo := repositories.NewUserRepo(dbPg)

	user := &entitiies.UserModel{
		Username: "admin",
		Password: "admin",
	}

	userRes, err := userRepo.InsertUser(user)

	if err != nil {
		t.Fatal(err.Error())
	}

	log.Printf("%#+v", userRes)

}
