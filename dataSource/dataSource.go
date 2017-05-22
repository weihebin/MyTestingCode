package dataSource

import (
	"MyTestingCode/localStorage"

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
	GetLastLocalAvailableDate() time.Time
}
type IDataProvider interface {
	RegisterColumn(column IColumn) error
	Subscribe(ctx context.Context, onReceived func(dataTypes.IData)) error
}
type dataProvider struct {
	symbol     string
	timeRange  TimeRange
	columns    map[string]IColumn
	onReceived func(dataTypes.IData)
	localCache localStorage.ILocalStorage
}

type TimeRange struct {
	sDate time.Time
	eDate time.Time
}

func (dp *dataProvider) RegisterColumn(column IColumn) error {
	requiredColumns := column.GetRequiredColumn()
	for idx := range requiredColumns {
		dp.RegisterColumn(requiredColumns[idx])
	}
	if nil == dp.columns {
		dp.columns = make(map[string]IColumn)
	}
	dp.columns[column.GetName()] = column
	return nil
}
func (dp *dataProvider) Subscribe(ctx context.Context, onData func(dataTypes.IData)) error {
	dp.onReceived = onData
	go dp.startPumping(ctx)
	return nil
}
func (dp *dataProvider) startPumping(ctx context.Context) {
	startDT := dp.getRemoteStartDate()
	dp.localCache.SetTimeRange(startDT, dp.timeRange.eDate)
	dp.localCache.LoadData(dp.symbol)

}
func (dp *dataProvider) getRemoteStartDate() time.Time {
	result := dp.timeRange.sDate
	for idx := range dp.columns {
		lastAvaliableDT := dp.columns[idx].GetLastLocalAvailableDate()
		if result.Sub(lastAvaliableDT).Seconds() > 0 {
			result = lastAvaliableDT
		}
	}
	return result
}
func NewDataProvider(symbol string, timeRange *TimeRange) dataTypes.IDataProvider {
	return &dataProvider{symbol: symbol, timeRange: *timeRange, localCache: localStorage.NewLocalStorage(symbol)}
}
