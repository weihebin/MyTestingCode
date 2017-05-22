package patternFinder

import "sync"
import "time"
import "encoding/json"

const (
	maxConcurrent = 8
)

type patternFinder struct {
	SearchingRange *searchingRangeWrapper
	Patterns       []*PatternWrapper
}

type SymbolData struct {
	Symbol    string
	PriceData []dataTypes.IData
}

func NewPatternFinder(params []byte) *patternFinder {
	var pf patternFinder

	json.Unmarshal(params, &pf)
	return &pf
}
func (pf *patternFinder) getSearchingRange() ISearchingRange {

	return pf.SearchingRange
}
func (pf *patternFinder) scanPatterns(symbol string, data []dataTypes.IData) {
	dr, _ := time.ParseDuration("10ms")
	time.Sleep(dr)
}
func (pf *patternFinder) workingRoutine(dataChannel chan SymbolData) {
	for sData := range dataChannel {
		pf.scanPatterns(sData.Symbol, sData.PriceData)
	}

}
func (pf *patternFinder) getSymbolTimeRange(symbol string, tr *TimeRange) *TimeRange {
	return tr
}
func (pf *patternFinder) getPriceDataOf(symbol string, tr *TimeRange) []dataTypes.IData {
	dr, _ := time.ParseDuration("1ms")
	time.Sleep(dr)
	return nil
}
func (pf *patternFinder) StartSearch() error {
	r := pf.getSearchingRange()
	symbols := r.GetSymbols()
	sCount := len(symbols)
	maxGroutines := maxConcurrent
	if maxGroutines > sCount {
		maxGroutines = sCount
	}
	wg := sync.WaitGroup{}
	dataChannel := make(chan SymbolData, maxGroutines)
	for n := 0; n < maxGroutines; n++ {

		wg.Add(1)
		go func() {
			pf.workingRoutine(dataChannel)
			wg.Done()
		}()
	}

	wg.Add(1)
	go func() {
		for idx := range symbols {
			sTimeRange := pf.getSymbolTimeRange(symbols[idx], r.GetTimeRange())
			data := pf.getPriceDataOf(symbols[idx], sTimeRange)
			dataChannel <- SymbolData{Symbol: symbols[idx], PriceData: data}
		}
		close(dataChannel)
		wg.Done()
	}()
	defer wg.Wait()

	return nil
}
func (pf *patternFinder) StartSearchSync() error {
	r := pf.getSearchingRange()
	symbols := r.GetSymbols()
	for idx := range symbols {

		sTimeRange := pf.getSymbolTimeRange(symbols[idx], r.GetTimeRange())
		data := pf.getPriceDataOf(symbols[idx], sTimeRange)
		pf.scanPatterns("", data)
	}
	return nil
}
