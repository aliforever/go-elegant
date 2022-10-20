package elegant

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

type scanner interface {
	Scan(dest ...any) error
}

type query[T tbl] struct {
	db        *sql.DB
	tableName string

	queryBuilder *QueryBuilder
}

func newQuery[T tbl](db *sql.DB, fn func(builder *QueryBuilder)) *query[T] {
	var t T

	q := newBuilder(t.TableName())

	if fn != nil {
		fn(q)
	}

	return &query[T]{
		db:           db,
		tableName:    t.TableName(),
		queryBuilder: q,
	}
}

func (q *query[T]) FindOne() (data *T, err error) {
	fields, err := q.fields()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", strings.Join(fields, ","), q.tableName, q.queryBuilder.builder.String())
	r := q.db.QueryRow(query, q.queryBuilder.values...)
	if r.Err() != nil {
		return nil, r.Err()
	}

	return q.decodeResult(fields, r)
}

func (q *query[T]) Find() (data []T, err error) {
	fields, err := q.fields()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(fields, ","), q.tableName)

	if q.queryBuilder.builder != nil && q.queryBuilder.builder.String() != "" {
		query += fmt.Sprintf(" WHERE %s", q.queryBuilder.builder.String())
	}

	cur, err := q.db.Query(query, q.queryBuilder.values...)
	if err != nil {
		return nil, err
	}
	defer cur.Close()

	for cur.Next() {
		if d, err := q.decodeResult(fields, cur); err != nil {
			continue
		} else {
			data = append(data, *d)
		}
	}

	return data, cur.Err()
}

func (q *query[T]) decodeResult(fields []string, sc scanner) (*T, error) {
	values := make([]interface{}, len(fields))

	for i := range fields {
		values[i] = &values[i]
	}

	err := sc.Scan(values...)
	if err != nil {
		return nil, err
	}

	d := map[string]interface{}{}
	for i, value := range values {
		d[fields[i]] = value
	}

	j, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	var data *T

	err = json.Unmarshal(j, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (q *query[T]) fields() (fields []string, err error) {
	var t T

	m, err := dataToMap(t)
	if err != nil {
		return nil, err
	}

	for key := range m {
		fields = append(fields, key)
	}

	return fields, nil
}
