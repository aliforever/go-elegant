package elegant

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aliforever/go-elegant/options"
	"strings"
)

type tbl interface {
	TableName() string
}

type Tbl[T tbl] struct {
	name    string
	db      *sql.DB
	options *options.TableOptions
}

func Table[T tbl](db *sql.DB, options ...*options.TableOptions) *Tbl[T] {
	var t T

	tbl := &Tbl[T]{name: t.TableName(), db: db}

	if len(options) != 0 {
		tbl.options = options[0]
	}

	return tbl
}

func (c *Tbl[T]) BuildSchema() *schema {
	return newSchemaBuilder(c.db, c.name)
}

func (c *Tbl[T]) Insert(data T, opts ...*options.InsertOptions) (*T, error) {
	m, err := dataToMap(data)
	if err != nil {
		return nil, err
	}

	ignoredFields := []string{}

	var insOptions *options.InsertOptions
	if len(opts) > 0 {
		insOptions = opts[0]
	} else if c.options != nil {
		insOptions = c.options.InsertOptions
	}

	if insOptions != nil {
		ignoredFields = opts[0].IgnoredFields
	}

	columnNames, placeHolders, values := c.mapToInsertQuery(m, ignoredFields...)

	var t T

	// TODO: Returning Id is only valid for postgresql
	//		This should be changed
	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s) RETURNING id`, t.TableName(), columnNames, placeHolders)

	var id interface{}
	err = c.db.QueryRow(query, values...).Scan(&id)
	if err != nil {
		return nil, err
	}

	j, err := json.Marshal(map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(j, &data)

	return &data, err
}

func (c *Tbl[T]) Query(fn func(builder *Builder)) *query[T] {
	return newQuery[T](c.db, fn)
}

// DropTable TODO: Separate to its own builder because it has many attributes
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

func (c *Tbl[T]) DropTableIfExistsCascade() (err error) {
	var t T
	str := fmt.Sprintf("DROP TABLE %s IF EXISTS CASCADE", t.TableName())
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
