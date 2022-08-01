package columns

import "fmt"

type boolean struct{}

func NewBoolean(name string) *Column[boolean] {
	return newColumn[boolean](name, boolean{})
}

func (c boolean) Builder() string {
	return fmt.Sprintf("%s", TypeBool)
}
