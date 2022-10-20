package elegant

import (
	"fmt"
	"strings"
)

type WhereClause struct {
	builder       *strings.Builder
	placeHolderFn func() int
	addValueFn    func(interface{})
	field         string
	operand       string
	val           interface{}
}

func newWhereClause(builder *strings.Builder, placeHolderFn func() int, addValueFn func(val interface{}), field, operand string, val interface{}) *WhereClause {
	_, _ = builder.WriteString(fmt.Sprintf("%s %s", field, operand))
	if placeHolderFn != nil {
		_, _ = builder.WriteString(fmt.Sprintf(" $%d", placeHolderFn()))
	}

	return &WhereClause{builder: builder, field: field, operand: operand, val: val, addValueFn: addValueFn, placeHolderFn: placeHolderFn}
}

func (w *WhereClause) And(field, operand string, val interface{}) *WhereClause {
	_, _ = w.builder.WriteString(fmt.Sprintf(" AND %s %s", field, operand))

	if w.placeHolderFn != nil {
		_, _ = w.builder.WriteString(fmt.Sprintf(" $%d", w.placeHolderFn()))
	}

	w.addValueFn(val)

	return w
}

func (w *WhereClause) Or(field, operand string, val interface{}) *WhereClause {
	_, _ = w.builder.WriteString(fmt.Sprintf(" OR %s %s", field, operand))

	if w.placeHolderFn != nil {
		_, _ = w.builder.WriteString(fmt.Sprintf(" $%d", w.placeHolderFn()))
	}

	w.addValueFn(val)

	return w
}

func (w *WhereClause) OrGroup(field, operand string, val interface{}, fn func(qb *WhereClause)) *WhereClause {
	w.group("OR", field, operand, val, fn)
	return w
}

func (w *WhereClause) AndGroup(field, operand string, val interface{}, fn func(qb *WhereClause)) *WhereClause {
	w.group("AND", field, operand, val, fn)
	return w
}

func (w *WhereClause) group(keyword, field, operand string, val interface{}, fn func(qb *WhereClause)) *WhereClause {
	w.builder.WriteString(fmt.Sprintf(" %s (", keyword))

	wc := newWhereClause(w.builder, w.placeHolderFn, w.addValueFn, field, operand, val)

	fn(wc)

	wc.builder.WriteString(")")

	w.addValueFn(val)

	return w
}
