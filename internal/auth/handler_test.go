package auth_test

import (
	"bytes"
	"encoding/json"
	"http_5/configs"
	"http_5/internal/auth"
	"http_5/internal/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"http_5/pkg/db"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}
	userRepo := user.NewUserRepository(&db.Db{
		DB: gormDb,
	})

	// Create full handler withou NewHandler
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}
	return &handler, mock, nil
}

// ex. of mockDB
func TestLoginHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("abc@a.com", "$2a$10$kbSAmj.MD.X5gw57Y7vCkutpSAf4tvm6MSPLRJXUR1BuUt4Mj1dHC")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	if err != nil {
		t.Fatal(err)
		return
	}
	// request response logic without server
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "abc@a.com",
		Password: "hello",
	})
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("got %d, expected %d", w.Code, 200)
	}
}

func TestRegisterHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password", "name"})

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	// "INSERT" can be more specific command
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()
	if err != nil {
		t.Fatal(err)
		return
	}
	// request response logic without server
	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "abc@a.com",
		Password: "hello",
		Name:     "TEST",
	})
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("got %d, expected %d", w.Code, 201)
	}
}
