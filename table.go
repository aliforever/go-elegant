package elegant

import (
	"database/sql"
	"fmt"
	"github.com/aliforever/go-elegant/options"
	"strings"
)

type tbl interface {
	TableName() string
}

type Tbl[T tbl] struct {
	name string
	db   *sql.DB
}

func Table[T tbl](db *sql.DB) *Tbl[T] {
	var t T

	return &Tbl[T]{name: t.TableName(), db: db}
}

func (c *Tbl[T]) BuildSchema() *schema {
	return newSchemaBuilder(c.db, c.name)
}

func (c *Tbl[T]) Insert(data T, opts ...*options.Insert) error {
	m, err := dataToMap(data)
	if err != nil {
		return err
	}

	ignoredFields := []string{}
	if len(opts) > 0 {
		ignoredFields = opts[0].IgnoredFields
	}

	columnNames, placeHolders, values := c.mapToInsertQuery(m, ignoredFields...)

	var t T

	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, t.TableName(), columnNames, placeHolders)

	_, err = c.db.Exec(query, values...)

	return err
}

func (c *Tbl[T]) Query(fn func(builder *Builder)) *query[T] {
	return newQuery[T](c.db, fn)
}

func (c *Tbl[T]) DropTable() (err error) {
	var t T
	str := fmt.Sprintf("DROP TABLE %s", t.TableName())
	_, err = c.db.Exec(str)
	return err
}

func (c *Tbl[T]) DropTableIfExists() (err error) {
	var t T
	str := fmt.Sprintf("DROP TABLE %s IF EXISTS", t.TableName())
	_, err = c.db.Exec(str)
	return err
}

func (c *Tbl[T]) mapToInsertQuery(m map[string]interface{}, ignoredFields ...string) (columnNames string, placeHolders string, values []interface{}) {
	isFieldIgnored := func(theField string) bool {
		if len(ignoredFields) == 0 {
			return false
		}
		for _, field := range ignoredFields {
			if strings.ToLower(field) == strings.ToLower(theField) {
				return true
			}
		}
		return false
	}

	cNames := []string{}

	for key, val := range m {
		if isFieldIgnored(key) {
			continue
		}
		cNames = append(cNames, key)
		values = append(values, val)
	}

	ps := []string{}
	for i := range values {
		ps = append(ps, fmt.Sprintf("$%d", i+1))
	}

	return strings.Join(cNames, ","), strings.Join(ps, ","), values
}
