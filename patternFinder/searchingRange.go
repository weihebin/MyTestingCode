package patternFinder

import (
	"encoding/json"
	"errors"
	"reflect"
)

type searchingRangeWrapper struct {
	searchingRange ISearchingRange
}

func (sr *searchingRangeWrapper) GetSymbols() []string {
	return sr.searchingRange.GetSymbols()
}

func (sr *searchingRangeWrapper) GetTimeRange() *TimeRange {
	return sr.searchingRange.GetTimeRange()
}

type testingSearchingRange struct {
	Symbols   []string
	TimeRange TimeRange
}

func (sr *testingSearchingRange) GetSymbols() []string {
	return sr.Symbols
}

func (sr *testingSearchingRange) GetTimeRange() *TimeRange {
	return &sr.TimeRange
}

func newSearchingRange(typeName string) (ret ISearchingRange, err error) {
	switch typeName {
	case "*patternFinder.testingSearchingRange":
		ret = &testingSearchingRange{}
	default:
		err = errors.New("type " + typeName + " is not supported")
	}
	return ret, err
}

func (p *searchingRangeWrapper) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(p.searchingRange)
	if nil == err {
		data, err = json.Marshal(
			struct {
				TypeName string
				Data     json.RawMessage
			}{
				TypeName: reflect.TypeOf(p.searchingRange).String(),
				Data:     json.RawMessage(data),
			})
	}
	return data, err
}

func (p *searchingRangeWrapper) UnmarshalJSON(data []byte) error {
	var tmpObj struct {
		TypeName string
		Data     json.RawMessage
	}

	err := json.Unmarshal(data, &tmpObj)
	if nil == err {
		obj, err := newSearchingRange(tmpObj.TypeName)
		if nil != obj && nil == err {
			json.Unmarshal(tmpObj.Data, obj)
			p.searchingRange = obj
		}
	}
	return err
}
