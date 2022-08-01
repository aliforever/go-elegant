package columns

import "fmt"

type bigint struct {
}

func NewBigInt(name string) *Column[bigint] {
	return newColumn[bigint](name, bigint{})
}

func (c bigint) Builder() string {
	return fmt.Sprintf("%s", TypeBigInt)
}
