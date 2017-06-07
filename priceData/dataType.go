package priceData

import (
	"time"
)

type IData interface {
}

type TimeRange struct {
	SDate time.Time
	EDate time.Time
}

func NewData() IData {
	return nil
}
