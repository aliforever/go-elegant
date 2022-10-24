package options

type TableOptions struct {
	PrimaryIDColumnName string
	InsertOptions       *InsertOptions
	UpdateOptions       *UpdateOptions
}

func Table() *TableOptions {
	return &TableOptions{}
}

func (t *TableOptions) SetInsertOptions(options *InsertOptions) *TableOptions {
	t.InsertOptions = options

	return t
}

func (t *TableOptions) SetUpdateOptions(options *UpdateOptions) *TableOptions {
	t.UpdateOptions = options

	return t
}

func (t *TableOptions) SetPrimaryIDColumnName(name string) *TableOptions {
	t.PrimaryIDColumnName = name
	return t
}
