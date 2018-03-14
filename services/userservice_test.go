package services

import (
	"reflect"
	"testing"

	"github.com/ducngocvnk57/circleci-demo/models"
)

func TestGetAllUser(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.User
		wantErr bool
	}{
		{
			name:    "test",
			want:    []models.User{models.User{}},
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
