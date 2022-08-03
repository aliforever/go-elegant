package elegant

import (
	"strings"
	"sync"
)

type Builder struct {
	tableName string

	builder *strings.Builder

	placeHolderIndex      int
	placeHolderIndexMutex sync.Mutex

	values []interface{}
}

func newBuilder(tblName string) *Builder {
	return &Builder{
		builder:          &strings.Builder{},
		placeHolderIndex: 0,
		tableName:        tblName,
	}
}

func (q *Builder) Where(fieldName, operand string, value interface{}) *whereClause {
	q.values = append(q.values, value)
	return newWhereClause(q.builder, q.newPlaceHolder, q.addValue, fieldName, operand, value)
}

func (q *Builder) addValue(val interface{}) {
	q.values = append(q.values, val)
}

func (q *Builder) newPlaceHolder() int {
	q.placeHolderIndexMutex.Lock()
	defer q.placeHolderIndexMutex.Unlock()

	q.placeHolderIndex++

	return q.placeHolderIndex
}
