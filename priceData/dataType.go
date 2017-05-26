package priceData

type IData interface {
}

type DataValueType string

const (
	DataValueType_Decimal DataValueType = "Decimal"
	DataValueType_String                = "String"
)

func NewData() IData {
	return nil
}
