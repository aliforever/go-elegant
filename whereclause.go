package elegant

import (
	"fmt"
	"strings"
)

type whereClause struct {
	builder       *strings.Builder
	placeHolderFn func() int
	addValueFn    func(interface{})
	field         string
	operand       string
	val           interface{}
}

func newWhereClause(builder *strings.Builder, placeHolderFn func() int, addValueFn func(val interface{}), field, operand string, val interface{}) *whereClause {
	_, _ = builder.WriteString(fmt.Sprintf("%s %s $%d", field, operand, placeHolderFn()))
	return &whereClause{builder: builder, field: field, operand: operand, val: val, addValueFn: addValueFn}
}

func (w *whereClause) And(field, operand string, val interface{}) *whereClause {
	_, _ = w.builder.WriteString(fmt.Sprintf(" AND %s %s $%d", field, operand, w.placeHolderFn()))
	w.addValueFn(val)
	return w
}

func (w *whereClause) Or(field, operand string, val interface{}) *whereClause {
	_, _ = w.builder.WriteString(fmt.Sprintf(" OR %s %s $%d", field, operand, w.placeHolderFn()))
	w.addValueFn(val)
	return w
}

func (w *whereClause) OrGroup(field, operand string, val interface{}, fn func(qb *whereClause)) *whereClause {
	w.group("OR", field, operand, val, fn)
	return w
}

func (w *whereClause) AndGroup(field, operand string, val interface{}, fn func(qb *whereClause)) *whereClause {
	w.group("AND", field, operand, val, fn)
	return w
}

func (w *whereClause) group(keyword, field, operand string, val interface{}, fn func(qb *whereClause)) *whereClause {
	w.builder.WriteString(fmt.Sprintf(" %s (", keyword))

	wc := newWhereClause(w.builder, w.placeHolderFn, w.addValueFn, field, operand, val)

	fn(wc)

	wc.builder.WriteString(")")

	w.addValueFn(val)

	return w
}
