package elegant

import (
	"strings"
	"sync"
)

type QueryBuilder struct {
	tableName string

	builder *strings.Builder

	placeHolderIndex      int
	placeHolderIndexMutex sync.Mutex

	values []interface{}
}

func newBuilder(tblName string) *QueryBuilder {
	return &QueryBuilder{
		builder:          &strings.Builder{},
		placeHolderIndex: 0,
		tableName:        tblName,
	}
}

func (q *QueryBuilder) Where(fieldName, operand string, value interface{}) *WhereClause {
	q.values = append(q.values, value)
	return newWhereClause(q.builder, q.newPlaceHolder, q.addValue, fieldName, operand, value)
}

func (q *QueryBuilder) addValue(val interface{}) {
	q.values = append(q.values, val)
}

func (q *QueryBuilder) newPlaceHolder() int {
	q.placeHolderIndexMutex.Lock()
	defer q.placeHolderIndexMutex.Unlock()

	q.placeHolderIndex++

	return q.placeHolderIndex
}
