package columns

import "fmt"

type realFloat struct{}

func NewReal(name string) *Column[realFloat] {
	return newColumn[realFloat](name, realFloat{})
}

func (c realFloat) Builder() string {
	return fmt.Sprintf("%s", TypeReal)
}
