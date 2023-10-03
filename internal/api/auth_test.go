package api

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/omerberkcan/banking-transfer/internal/repository"
	"github.com/omerberkcan/banking-transfer/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbSetup() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return gormDB, mock
}

func apiSetup(db *gorm.DB) *Handlers {
	repo := repository.New(db)
	s := service.New(repo)
	handlers := NewHandler(s)
	return handlers
}

func TestRegister(t *testing.T) {
	// t.Run("should return 201 status ok", func(t *testing.T) {
	// 	DB, mock := dbSetup()
	// 	handler := apiSetup(DB)

	// 	userJSON := `{"id_no":"12345678911","name":"Omer","balance":10.5,"password":"12345"}`
	// 	e := NewEcho()
	// 	req := httptest.NewRequest(http.MethodPost, "/v1/register", strings.NewReader(userJSON))
	// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// 	rec := httptest.NewRecorder()
	// 	c := e.NewContext(req, rec)

	// 	expectedQuery := `INSERT INTO "users" ("created_at","updated_at","deleted_at","id_no","name","password","balance") VALUES (?,?,?,?,?,?,?)`

	// 	mock.ExpectExec(regexp.QuoteMeta(expectedQuery)).
	// 		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "12345678911", "Omer", "12345", 10.5).
	// 		WillReturnResult(sqlmock.NewResult(1, 1))
	// 	mock.ExpectCommit()
	// 	if err := mock.ExpectationsWereMet(); err != nil {
	// 		t.Fatalf("Unfulfilled expectations: %v", err)
	// 	}
	// 	if assert.NoError(t, handler.Auth.Register(c)) {
	// 		assert.Equal(t, http.StatusCreated, rec.Code)
	// 	}
	// })
}
