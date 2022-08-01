package columns

import "fmt"

type smallint struct {
}

func NewSmallInt(name string) *Column[smallint] {
	return newColumn[smallint](name, smallint{})
}

func (c smallint) Builder() string {
	return fmt.Sprintf("%s", TypeSmallInt)
}
