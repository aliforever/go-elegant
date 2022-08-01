package columns

type dataType string

const (
	TypeBool    dataType = "BOOLEAN"
	TypeChar    dataType = "CHAR"
	TypeVarchar dataType = "VARCHAR"
	TypeText    dataType = "TEXT"

	TypeSmallInt dataType = "SMALLINT" // 2 bytes [-32768 to +32767]
	TypeInt      dataType = "INTEGER"  // 4 bytes [-2147483648 to +2147483647]
	TypeBigInt   dataType = "BIGINT"   // 8 bytes [-9223372036854775808 to +9223372036854775807]

	TypeReal   dataType = "REAL"             //	4 bytes [6 decimal digits precision]
	TypeDouble dataType = "DOUBLE PRECISION" // 8 bytes [15 decimal digits precision]
)
