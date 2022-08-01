package columns

import "fmt"

type text struct {
}

func NewText(name string) *Column[text] {
	return newColumn[text](name, text{})
}

func (c text) Builder() string {
	return fmt.Sprintf("%s", TypeText)
}
