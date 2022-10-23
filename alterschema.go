package elegant

import (
	"database/sql"
	"fmt"
	"github.com/aliforever/go-elegant/columns"
	"strings"
)

type AlterSchema struct {
	db        *sql.DB
	tableName string

	columns     []columns.DataType
	foreignKeys []string
}

func newAlterSchemaBuilder(db *sql.DB, tableName string) *AlterSchema {
	return &AlterSchema{db: db, tableName: tableName}
}

func (s *AlterSchema) AddColumn(column columns.DataType) *AlterSchema {
	s.columns = append(s.columns, column)

	return s
}

func (s *AlterSchema) AddForeignKey(key string, targetTableName string, targetKeyName string) *AlterSchema {
	s.foreignKeys = append(s.foreignKeys, fmt.Sprintf("Add FOREIGN KEY (%s) REFERENCES %s(%s)", key, targetTableName, targetKeyName))

	return s
}

func (s *AlterSchema) tableData() string {
	var strs []string

	for _, column := range s.columns {
		strs = append(strs, fmt.Sprintf("Add %s", column.Builder()))
	}

	strs = append(strs, s.foreignKeys...)

	return strings.Join(strs, ",")
}

func (s *AlterSchema) buildQuery() string {
	return fmt.Sprintf("ALTER TABLE %s %s", s.tableName, s.tableData())
}

func (s *AlterSchema) Build() error {
	_, err := s.db.Exec(s.buildQuery())
	return err
}
