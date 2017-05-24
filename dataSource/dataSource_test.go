package dataSource

import (
	"MyTestingCode/dataTypes"
	"context"
	"reflect"
	"testing"
	"time"
)

func Test_dataProvider_RegisterColumn(t *testing.T) {
	dp := NewDataProvider("ABC", &TimeRange{sDate: time.Now().AddDate(-1, 0, 0), eDate: time.Now()})
	tc := &testColumn{name: "ABC"}
	type args struct {
		column IColumn
	}
	tests := []struct {
		name    string
		dp      IDataProvider
		args    args
		wantErr bool
	}{
		{
			name: "Register Simple column",
			dp:   dp,
			args: args{column: tc},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.dp.RegisterColumn(tt.args.column); (err != nil) != tt.wantErr {
				t.Errorf("dataProvider.RegisterColumn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestColumnDependencies(t *testing.T) {
	// read data
}

type testColumn struct {
	name string
}

func (c *testColumn) GetRequiredColumn() []IColumn {
	return nil
}
func (c *testColumn) GetName() string {
	return c.name
}
func (c *testColumn) GetBasicDataStartDate() time.Time {
	return time.Now()
}
func (c *testColumn) CalculateValue(data dataTypes.IData) {

}

func BenchmarkRegisterColumn(b *testing.B) {
	dp := NewDataProvider("ABC", &TimeRange{sDate: time.Now().AddDate(-1, 0, 0), eDate: time.Now()})
	tc := &testColumn{name: "ABC"}
	dp.RegisterColumn(tc)
}

func Test_dataProvider_Subscribe(t *testing.T) {
	type args struct {
		ctx    context.Context
		onData func(dataTypes.IData)
	}
	tests := []struct {
		name    string
		dp      *dataProvider
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.dp.Subscribe(tt.args.ctx, tt.args.onData); (err != nil) != tt.wantErr {
				t.Errorf("dataProvider.Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dataProvider_getColumnsInDependencyOrder(t *testing.T) {
	tests := []struct {
		name string
		dp   *dataProvider
		want []IColumn
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dp.getColumnsInDependencyOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataProvider.getColumnsInDependencyOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
