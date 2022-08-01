package columns

import "fmt"

type Integer struct {
}

func NewInteger(name string) *Column[Integer] {
	return newColumn[Integer](name, Integer{})
}

func (c Integer) Builder() string {
	return fmt.Sprintf("%s", TypeInt)
}
