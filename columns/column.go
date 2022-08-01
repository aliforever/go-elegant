package columns

import "fmt"

type ColumnInterface interface {
	Query() string
}

type DataType interface {
	Builder() string
}

type Column[T DataType] struct {
	data           T
	name           string
	isPK           bool
	isIdentity     bool
	notNull        bool
	unique         bool
	referredTable  string
	referredColumn string
}

func newColumn[T DataType](name string, t T) *Column[T] {
	return &Column[T]{name: name, data: t}
}

func (c *Column[T]) PrimaryKey() *Column[T] {
	c.isPK = true
	return c
}

func (c *Column[T]) Identity() *Column[T] {
	c.isIdentity = true
	return c
}

func (c *Column[T]) NotNull() *Column[T] {
	c.notNull = true
	return c
}

func (c *Column[T]) Unique() *Column[T] {
	c.unique = true
	return c
}

func (c *Column[T]) ForeignKey(sourceTable, column string) *Column[T] {
	c.referredTable = sourceTable
	c.referredColumn = column
	return c
}

func (c Column[T]) Builder() string {
	query := fmt.Sprintf("%s %s", c.name, c.data.Builder())
	if c.isPK {
		query += " PRIMARY KEY"
	}
	if c.notNull {
		query += " NOT NULL"
	}
	if c.unique {
		query += " UNIQUE"
	}
	if c.isIdentity {
		query += " GENERATED ALWAYS AS IDENTITY"
	}

	return query
}
