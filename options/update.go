package options

type UpdateOptions struct {
	IgnoredFields []string
}

func Update() *UpdateOptions {
	return &UpdateOptions{}
}

func (i *UpdateOptions) IgnoreFields(fields ...string) *UpdateOptions {
	i.IgnoredFields = append(i.IgnoredFields, fields...)
	return i
}
