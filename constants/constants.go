package constants

type ColumnCatalog string
type DataValueType uint

const (
	DataValueType_Int DataValueType = iota + 1
	DataValueType_String
	DataValueType_Decimal
	ColumnCatalog_Basic ColumnCatalog = "Basic"
)