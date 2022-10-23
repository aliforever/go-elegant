package elegant

import (
	"database/sql"
	"fmt"
	"github.com/aliforever/go-elegant/columns"
	"strings"
)

type buildSchema struct {
	db        *sql.DB
	tableName string

	columns      []columns.DataType
	dropIfExists bool
}

func newBuildSchemaBuilder(db *sql.DB, tableName string) *buildSchema {
	return &buildSchema{db: db, tableName: tableName}
}

func (s *buildSchema) DropTableIfExist() *buildSchema {
	s.dropIfExists = true

	return s
}

func (s *buildSchema) AddColumn(column columns.DataType) *buildSchema {
	s.columns = append(s.columns, column)

	return s
}

func (s *buildSchema) tableData() string {
	var strs []string

	for _, column := range s.columns {
		strs = append(strs, column.Builder())
	}

	return strings.Join(strs, ",")
}

func (s *buildSchema) Build() error {
	str := fmt.Sprintf("CREATE TABLE %s (%s)", s.tableName, s.tableData())
	if s.dropIfExists {
		str = fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE; ", s.tableName) + str
	}
	_, err := s.db.Exec(str)
	return err
}
