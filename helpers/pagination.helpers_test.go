package helpers

import (
	"reflect"
	"skeleton-svc/helpers/models"
	"testing"
)

func TestPagination(t *testing.T) {
	type args struct {
		page      int
		limit     int
		totalData int64
	}
	tests := []struct {
		name string
		args args
		want *models.Page
	}{
		{
			name: "Test Pagination",
			args: args{
				page:      1,
				limit:     10,
				totalData: 100,
			},
			want: &models.Page{
				CurrentPage:  1,
				PreviousPage: 0,
				NextPage:     2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pagination(tt.args.page, tt.args.limit, tt.args.totalData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pagination() = %v, want %v", got, tt.want)
			}
		})
	}
}
