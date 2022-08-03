package columns

import "fmt"

type timestamptz struct {
}

func NewTimestampWithTimezone(name string) *Column[timestamptz] {
	return newColumn[timestamptz](name, timestamptz{})
}

func (c timestamptz) Builder() string {
	return fmt.Sprintf("%s", TypeTimestampWithTimezone)
}
