package options

type InsertOptions struct {
	IgnoredFields []string
}

func Insert() *InsertOptions {
	return &InsertOptions{}
}

func (i *InsertOptions) IgnoreFields(fields ...string) *InsertOptions {
	i.IgnoredFields = append(i.IgnoredFields, fields...)
	return i
}
