package patternFinder

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshal_Unmarshal_Patterns(t *testing.T) {
	p, _ := NewPattern("*patternFinder.TestPattern")
	assert.NotNil(t, p)
	ps := make([]*PatternWrapper, 0)
	ps = append(ps, p)
	data, _ := json.Marshal(ps)
	assert.NotEmpty(t, data)
	var ps2 []*PatternWrapper

	json.Unmarshal(data, &ps2)

	assert.NotNil(t, ps2)
	assert.Equal(t, ps, ps2)

}

func BenchmarkMarshalAndUnmarshalPatterns(b *testing.B) {
	for n := 0; n < b.N; n++ {
		p, _ := NewPattern("*patternFinder.TestPattern")
		ps := make([]*PatternWrapper, 0)
		ps = append(ps, p)
		data, _ := json.Marshal(ps)

		var ps2 []*PatternWrapper

		json.Unmarshal(data, &ps2)
	}
}
