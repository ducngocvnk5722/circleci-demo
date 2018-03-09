package services

import (
	"reflect"
	"testing"

	"github.com/ducngocvnk57/circleci-demo/models"
)

func TestGetAllJobType(t *testing.T) {
	tests := []struct {
		name string
		want []models.JobType
	}{
		{
			name: "test",
			want: []models.JobType{},
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
