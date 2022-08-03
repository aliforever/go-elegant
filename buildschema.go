package elegant

import (
	"database/sql"
	"fmt"
	"github.com/aliforever/go-elegant/columns"
	"strings"
)

type schema struct {
	db        *sql.DB
	tableName string

	columns []columns.DataType
}

func newSchemaBuilder(db *sql.DB, tableName string) *schema {
	return &schema{db: db, tableName: tableName}
}

func (s *schema) AddColumn(column columns.DataType) *schema {
	s.columns = append(s.columns, column)

	return s
}

func (s *schema) tableData() string {
	var strs []string

	for _, column := range s.columns {
		strs = append(strs, column.Builder())
	}

	return strings.Join(strs, ",")
}

func (s *schema) Build() error {
	str := fmt.Sprintf("CREATE TABLE %s (%s)", s.tableName, s.tableData())
	_, err := s.db.Exec(str)
	return err
}
