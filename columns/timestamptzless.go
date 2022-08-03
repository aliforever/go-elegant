package columns

import "fmt"

type timestamptzless struct {
}

func NewTimestampWithoutTimezone(name string) *Column[timestamptzless] {
	return newColumn[timestamptzless](name, timestamptzless{})
}

func (c timestamptzless) Builder() string {
	return fmt.Sprintf("%s", TypeTimestampWithoutTimezone)
}
