package options

type TableOptions struct {
	InsertOptions *InsertOptions
}

func Table() *TableOptions {
	return &TableOptions{}
}

func (t *TableOptions) SetInsertOptions(options *InsertOptions) *TableOptions {
	t.InsertOptions = options

	return t
}
