package columns

import "fmt"

type Varchar struct {
	length int64
}

func NewVarchar(name string, length int64) *Column[Varchar] {
	return newColumn[Varchar](name, Varchar{length: length})
}

func (c Varchar) Builder() string {
	return fmt.Sprintf("%s (%d)", TypeVarchar, c.length)
}
