package elegant

import (
	"database/sql"
	"github.com/aliforever/go-elegant/columns"
	"testing"
)

func TestAlterSchema_buildQuery(t *testing.T) {
	type fields struct {
		db          *sql.DB
		tableName   string
		columns     []columns.DataType
		foreignKeys []struct {
			Key         string
			TargetTable string
			TargetKey   string
		}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Built",
			fields: fields{
				db:        nil,
				tableName: "users",
				columns: []columns.DataType{
					columns.NewInteger("id").Identity().NotNull(),
					columns.NewInteger("parent_id").NotNull(),
				},
				foreignKeys: []struct {
					Key         string
					TargetTable string
					TargetKey   string
				}{
					{
						Key:         "parent_id",
						TargetTable: "parent_users",
						TargetKey:   "id",
					},
				},
			},
			want: "ALTER TABLE users Add id INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,Add parent_id INTEGER NOT NULL,Add FOREIGN KEY (parent_id) REFERENCES parent_users(id)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newAlterSchemaBuilder(tt.fields.db, tt.fields.tableName)
			for _, column := range tt.fields.columns {
				s.AddColumn(column)
			}
			for _, key := range tt.fields.foreignKeys {
				s.AddForeignKey(key.Key, key.TargetTable, key.TargetKey)
			}
			got := s.buildQuery()

			if got != tt.want {
				t.Errorf("buildQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
