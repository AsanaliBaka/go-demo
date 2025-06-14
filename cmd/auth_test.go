package main

import (
	"bytes"
	"encoding/json"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/user"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "asan@brat.03",
		Password: "	$2a$10$/cZrcSOsAXFoDGeawHjMROK5fVXddVyY5FteZIQViz2CV.qQt5qIS",
		Name:     "brathu",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?,asan@brat.03 ").
		Delete(&user.User{})
}
func TestLoginSucces(t *testing.T) {
	//prepare
	db := initDb()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()
	defer removeData(db)
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "asan@brat.03	",
		Password: "123456Ba91",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatal(res)
	}

	result, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var resData auth.LoginResponse

	err = json.Unmarshal(result, &resData)
	if err != nil {
		t.Fatal(err)
	}

	if resData.Token == "" {
		t.Fatal("token 0")
	}

	t.Log(resData.Token)

}
