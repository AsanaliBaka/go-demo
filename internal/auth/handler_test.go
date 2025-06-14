package auth_test

import (
	"bytes"
	"encoding/json"
	conifgs "go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func boodcap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()

	if err != nil {
		return nil, nil, err
	}

	gorm, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))

	if err != nil {
		return nil, nil, err
	}

	userRepo := user.NewUserRepo(&db.Db{
		DB: gorm,
	})

	handler := auth.AuthHandler{
		Config: &conifgs.Config{
			Auth: conifgs.AuthConfig{
				Secret: "1234",
			},
		},

		AuthService: auth.NewAuthService(userRepo),
	}

	return &handler, mock, nil

}

func TestLoging(t *testing.T) {

	handler, mock, err := boodcap()

	rows := sqlmock.NewRows([]string{"email", "password"}).AddRow(
		"asan@brat.03", "$2a$10$/cZrcSOsAXFoDGeawHjMROK5fVXddVyY5FteZIQViz2CV.qQt5qIS")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if err != nil {
		t.Fatal(err)
	}

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "asan@brat.03	",
		Password: "123456B",
	})

	reader := bytes.NewReader(data)

	wr := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)

	handler.Login()(wr, req)

	if wr.Result().StatusCode != http.StatusOK {
		t.Error(wr.Code)
	}

}

func TestRegister(t *testing.T) {
	handler, mock, err := boodcap()

	if err != nil {
		t.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "asan@a.ru",
		Password: "1234awef",
		Name:     "asanali",
	})

	record := bytes.NewReader(data)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/register", record)

	handler.Register()(w, r)

	if w.Code != 200 {
		t.Fatal(w.Code)
	}

}
