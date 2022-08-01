package options

type Insert struct {
	IgnoredFields []string
}

func NewInsert() *Insert {
	return &Insert{}
}

func (i *Insert) IgnoreFields(fields ...string) *Insert {
	i.IgnoredFields = append(i.IgnoredFields, fields...)
	return i
}
