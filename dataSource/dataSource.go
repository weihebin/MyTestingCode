package dataSource

import (
	"MyTestingCode/localStorage"
	"container/list"

	"MyTestingCode/dataTypes"
	"context"
	"time"
)

type IColumnAccessor interface {
	GetValue(data dataTypes.IData) interface{}
}
type IColumn interface {
	GetRequiredColumn() []IColumn
	GetName() string
	//GetBasicDataStartDate return the earlies date required from basic columns in orde rgenerate the value for the given date
	GetBasicDataStartDate(date *time.Time) time.Time
	CalculateValue(data dataTypes.IData) error
}
type IDataProvider interface {
	RegisterColumn(column IColumn) error
	Subscribe(ctx context.Context, onReceived func(dataTypes.IData)) error
}
type dataProvider struct {
	symbol     string
	timeRange  TimeRange
	columns    []IColumn
	onReceived func(dataTypes.IData)
	localCache localStorage.ILocalStorage
}

type TimeRange struct {
	sDate time.Time
	eDate time.Time
}

func (dp *dataProvider) RegisterColumn(column IColumn) error {
	// requiredColumns := column.GetRequiredColumn()
	// for idx := range requiredColumns {
	// 	dp.RegisterColumn(requiredColumns[idx])
	// }
	if nil == dp.columns {
		dp.columns = make([]IColumn, 0, 1)
	}
	dp.columns = append(dp.columns, column)
	return nil
}
func (dp *dataProvider) Subscribe(ctx context.Context, onData func(dataTypes.IData)) error {
	dp.onReceived = onData
	go dp.startPumping(ctx)
	return nil
}

// return the slice in dependency order
func (dp *dataProvider) getColumnsInDependencyOrder() []IColumn {

	resolvedList := list.New()
	unresolvedList := list.New()
	for _, column := range dp.columns {
		depResolve(column, resolvedList, unresolvedList)
	}
	result := make([]IColumn, 0, resolvedList.Len())
	for e := resolvedList.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(IColumn))
	}
	return result
}

// https://www.electricmonk.nl/docs/dependency_resolving_algorithm/dependency_resolving_algorithm.html#_dependency_resolution_order
func depResolve(node IColumn, resolved *list.List, unresolved *list.List) {
	elementToRemove := unresolved.PushBack(node)
	dependencies := node.GetRequiredColumn()
	for _, edge := range dependencies {
		if !isColumnInTheList(edge, resolved) {
			if isColumnInTheList(edge, unresolved) {
				panic("circle dependency detected " + node.GetName() + "->" + edge.GetName())
			}
			depResolve(edge, resolved, unresolved)
		}
	}
	resolved.PushBack(node)
	unresolved.Remove(elementToRemove)
}
func isColumnInTheList(column IColumn, list *list.List) bool {
	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value.(IColumn).GetName() == column.GetName() {
			return true
		}
	}
	return false
}
func (dp *dataProvider) getBasicColumns(columns []IColumn) []IColumn {
	return nil
}
func (dp *dataProvider) getNotBasicColumns(columns []IColumn) []IColumn {
	return nil
}

// get data for basic columns, if data not cached, then retrieve from remove
// block while waiting remove
// return nil means not more data available
func (dp *dataProvider) getDataOfBasicColumns(startDT *time.Time, columns []IColumn) (dataTypes.IData, time.Time) {

	return nil, time.Time{}
}
func (dp *dataProvider) getValueFromLocal(data dataTypes.IData, column IColumn) bool {
	return false
}
func (dp *dataProvider) updateLocal(data dataTypes.IData) {

}
func (dp *dataProvider) startPumping(ctx context.Context) {
	columnsInOrder := dp.getColumnsInDependencyOrder()
	basicColumns := dp.getBasicColumns(columnsInOrder)
	nonBasicColumns := dp.getNotBasicColumns(columnsInOrder)
	startDT := dp.getEarliestStartDate()

	for data, startDT := dp.getDataOfBasicColumns(&startDT, basicColumns); nil != data; data, startDT = dp.getDataOfBasicColumns(&startDT, basicColumns) {
		bUpdateLocal := false
		for _, column := range nonBasicColumns {

			readFromLocal := dp.getValueFromLocal(data, column)
			if !readFromLocal { // not exists in local
				column.CalculateValue(data)
				bUpdateLocal = true
			}
		}
		if bUpdateLocal {
			dp.updateLocal(data)
		}
		if dp.timeRange.eDate.Sub(startDT).Hours() < 0 {
			break
		}
	}
	dp.localCache.SetTimeRange(startDT, dp.timeRange.eDate)
	dp.localCache.LoadData(dp.symbol)

}
func returnSmallerDate(date1, date2 *time.Time) *time.Time {
	if date1.Sub(*date2).Seconds() > 0 {
		return date2
	}
	return date1
}

// return the next date market open
func nextMarketDate(date *time.Time) time.Time {
	// TODO: not implemented ytet
	return *date
}
func (dp *dataProvider) getLocalAvailableDateForColumn(column IColumn) time.Time {
	return time.Time{}
}

func (dp *dataProvider) getEarliestStartDate() time.Time {
	result := dp.timeRange.sDate
	columnsInOrder := dp.getColumnsInDependencyOrder()
	for n := len(columnsInOrder); n >= 0; n-- {
		localAvailableDate := dp.getLocalAvailableDateForColumn(columnsInOrder[n])
		startDate := returnSmallerDate(&dp.timeRange.sDate, &localAvailableDate)
		result = columnsInOrder[n].GetBasicDataStartDate(startDate)
	}
	return result
}
func NewDataProvider(symbol string, timeRange *TimeRange) IDataProvider {
	return &dataProvider{symbol: symbol, timeRange: *timeRange, localCache: localStorage.NewLocalStorage(symbol)}
}
