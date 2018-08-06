package services

import (
	"reflect"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/ducngocvnk57/circleci-demo/models"
)

func TestGetAllJobType(t *testing.T) {
	jobtype := models.JobType{
		Id:        1,
		Title:     "unitest jobtype",
		CompanyId: 1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	query := regexp.QuoteMeta(`SELECT * FROM "job_types"`)
	rows := sqlmock.NewRows([]string{"id", "title", "company_id", "created_at", "updated_at"}).
		AddRow(jobtype.Id, jobtype.Title, jobtype.CompanyId, jobtype.CreatedAt, jobtype.UpdatedAt)
	mock.ExpectQuery(query).WillReturnRows(rows)
	tests := []struct {
		name string
		want []models.JobType
	}{
		{
			name: "test",
			want: []models.JobType{jobtype},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllJobType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllJobType() = %v, want %v", got, tt.want)
			}
		})
	}
}
