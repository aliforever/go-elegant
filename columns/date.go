package columns

import "fmt"

type date struct {
}

func NewDate(name string) *Column[date] {
	return newColumn[date](name, date{})
}

func (c date) Builder() string {
	return fmt.Sprintf("%s", TypeDate)
}
