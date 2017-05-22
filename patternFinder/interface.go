package patternFinder

import "time"

type IPatternFinder interface {
	StartSearch() error
}
type TimeRange struct {
	StartTime time.Time
	EndTime   time.Time
}
type ISearchingRange interface {
	GetSymbols() []string
	GetTimeRange() *TimeRange
}
type IPattern interface {
}
