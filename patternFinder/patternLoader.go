package patternFinder

import (
	"encoding/json"
	"errors"
	"reflect"
)

type PatternWrapper struct {
	pattern IPattern
}
type TestPattern struct {
}

func NewPattern(typeName string) (ret *PatternWrapper, err error) {
	p, err := newPattern(typeName)
	if nil != p && nil == err {
		return &PatternWrapper{
			pattern: p,
		}, err
	}
	return nil, err
}
func newPattern(typeName string) (ret IPattern, err error) {
	switch typeName {
	case "*patternFinder.TestPattern":
		ret = &TestPattern{}
	default:
		err = errors.New("type " + typeName + " is not supported")
	}
	return ret, err
}

func (p *PatternWrapper) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(p.pattern)
	if nil == err {
		data, err = json.Marshal(
			struct {
				TypeName string
				Data     json.RawMessage
			}{
				TypeName: reflect.TypeOf(p.pattern).String(),
				Data:     json.RawMessage(data),
			})
	}
	return data, err
}

func (p *PatternWrapper) UnmarshalJSON(data []byte) error {
	var tmpObj struct {
		TypeName string
		Data     json.RawMessage
	}

	err := json.Unmarshal(data, &tmpObj)
	if nil == err {
		obj, err := newPattern(tmpObj.TypeName)
		if nil != obj && nil == err {
			json.Unmarshal(tmpObj.Data, obj)
			p.pattern = obj
		}
	}
	return err
}
