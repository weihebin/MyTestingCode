package patternFinder

import (
	"testing"

	"sync"

	"time"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestPatternFinder(t *testing.T) {
	pf := NewPatternFinder([]byte("{\"SearchingRange\":{},\"Patterns\":[{\"TypeName\":\"*patternFinder.TestPattern\",\"Data\":{}}]}"))
	assert.NotNil(t, pf)

	pf.SearchingRange = getTestingRange()

	data, _ := json.Marshal(pf)
	println(string(data))

	var pf2 patternFinder

	json.Unmarshal(data, &pf2)
	assert.Equal(t, *pf, pf2)
	pf.StartSearch()

}
func getTestingRange() *searchingRangeWrapper {
	return &searchingRangeWrapper{
		searchingRange: &testingSearchingRange{
			Symbols: []string{"BABA", "NVDA"},
			TimeRange: TimeRange{
				StartTime: time.Now().AddDate(-1, 0, 0),
				EndTime:   time.Now(),
			},
		}}
}
func BenchmarkPatternFinderSync(b *testing.B) {

	for n := 0; n < b.N; n++ {
		pf := NewPatternFinder([]byte("{}"))
		pf.SearchingRange = getTestingRange()

		pf.StartSearchSync()
	}
}

func BenchmarkPatternFinderAsync(b *testing.B) {

	wg := sync.WaitGroup{}

	for n := 0; n < b.N; n++ {
		wg.Add(1)
		go func() {
			pf := NewPatternFinder([]byte("{}"))
			pf.SearchingRange = getTestingRange()

			pf.StartSearch()
			wg.Done()
		}()
	}
	wg.Wait()
}
