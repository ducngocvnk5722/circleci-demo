package services

import (
	"database/sql"
	"os"
	"reflect"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/ducngocvnk57/circleci-demo/models"
	"github.com/jinzhu/gorm"
)

var mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	var err error
	var mockDB *sql.DB
	mockDB, mock, err = sqlmock.NewWithDSN("sqlmock_db_company")
	if err != nil {
		panic("Got an unexpected error.")
	}
	db, err = gorm.Open("sqlmock", mockDB)
	if err != nil {
		panic("Got an unexpected error.")
	}
	defer func() {
		db.Close()
		mockDB.Close()
	}()
	os.Exit(m.Run())
}
func TestGetAllUser(t *testing.T) {
	query := regexp.QuoteMeta(`SELECT * FROM "users"`)
	user := models.User{
		Id:        1,
		Name:      "unitest",
		CompanyId: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	rows := sqlmock.NewRows([]string{"id", "name", "company_id", "created_at", "updated_at"}).
		AddRow(user.Id, user.Name, user.CompanyId, user.CreatedAt, user.UpdatedAt)
	mock.ExpectQuery(query).WillReturnRows(rows)
	tests := []struct {
		name    string
		want    []models.User
		wantErr bool
	}{
		{
			name:    "test",
			want:    []models.User{user},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
