package columns

import "fmt"

type doublePrecision struct{}

func NewDoublePrecision(name string) *Column[doublePrecision] {
	return newColumn[doublePrecision](name, doublePrecision{})
}

func (c doublePrecision) Builder() string {
	return fmt.Sprintf("%s", TypeDouble)
}
