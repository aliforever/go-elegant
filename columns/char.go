package columns

import "fmt"

type char struct {
	length int64
}

func NewChar(name string, length int64) *Column[char] {
	return newColumn[char](name, char{length: length})
}

func (c char) Builder() string {
	return fmt.Sprintf("%s(%d)", TypeChar, c.length)
}
