package elegant

type rawQueryBuilder struct {
	tableName string

	rawQuery string

	values []interface{}
}

func (q *rawQueryBuilder) Query() string {
	return q.rawQuery
}

func (q *rawQueryBuilder) Values() []interface{} {
	return q.values
}

func newRawBuilder(tblName string, q string, values []interface{}) *rawQueryBuilder {
	return &rawQueryBuilder{
		tableName: tblName,
		rawQuery:  q,
		values:    values,
	}
}
