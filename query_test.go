package elegant

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

type testTable struct {
}

func (t testTable) TableName() string {
	return "groups"
}

func Test_query_QueryRaw(t *testing.T) {
	type fields struct {
		db           *sql.DB
		tableName    string
		queryBuilder queryBuilder
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  []interface{}
	}{
		{
			name:   "Case 1",
			fields: fields{},
			want:   "(SELECT group_id FROM group_user WHERE user_id = $1)",
			want1:  []interface{}{"6c4283cf-cc95-41a4-937c-4b821c2da147"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := newRawQuery[testTable](nil, "SELECT group_id FROM group_user WHERE user_id = $1", "6c4283cf-cc95-41a4-937c-4b821c2da147")
			got, got1 := q.Query()
			fmt.Println(got, got1)
			if got != tt.want {
				t.Errorf("Query() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Query() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
