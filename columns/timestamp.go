package columns

import "fmt"

type timestamp struct {
}

func NewTimestamp(name string) *Column[timestamp] {
	return newColumn[timestamp](name, timestamp{})
}

func (c timestamp) Builder() string {
	return fmt.Sprintf("%s", TypeTimestamp)
}
