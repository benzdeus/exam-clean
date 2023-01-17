package db_test

import (
	"demo-clean/db"
	"testing"
)

func TestConnectDB(t *testing.T) {
	_, tearDown := db.New()

	defer tearDown()

}
