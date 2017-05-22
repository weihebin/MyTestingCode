package localStorage

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_localStorage_getLocation(t *testing.T) {
	ls := setupTestLocalStorage()
	os.Setenv("LOCAL_STORAGE_ROOT", "root")
	type args struct {
		symbol string
	}
	tests := []struct {
		name string
		ls   *localStorage
		args args
		want string
	}{
		{
			name: "GetLocation",
			ls:   ls,
			args: args{symbol: "abcd"},
			want: "root/a/abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ls.getLocation(tt.args.symbol); got != tt.want {
				t.Errorf("localStorage.getLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
func setupTestLocalStorage() *localStorage {
	os.Setenv("LOCAL_STORAGE_ROOT", "C:/Users/hwei/AppData/Local/Temp/")

	return &localStorage{}

}
func Test_localStorage_LoadData(t *testing.T) {

	ls := setupTestLocalStorage()
	type args struct {
		symbol string
	}
	tests := []struct {
		name string
		ls   *localStorage
		args args
	}{
		{
			name: "GetLocation",
			ls:   ls,
			args: args{symbol: "appl"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ls.LoadData(tt.args.symbol)
		})
	}
}

func Test_localStorage_createIfNotExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		ls   *localStorage
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ls.createIfNotExists(tt.args.path)
		})
	}
}

func Test_localStorage_GetAvailableConnections(t *testing.T) {
	tests := []struct {
		name string
		ls   *localStorage
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ls.GetAvailableConnections()
		})
	}
}

func Test_localStorage_SetMaxConnection(t *testing.T) {
	tests := []struct {
		name string
		ls   *localStorage
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ls.SetMaxConnection()
		})
	}
}

func Test_localStorage_startConnectionPool(t *testing.T) {
	tests := []struct {
		name string
		ls   *localStorage
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ls.startConnectionPool()
		})
	}
}

func Test_localStorage_RegisterCacheConsumer(t *testing.T) {
	type args struct {
		consumer ICacheDataConsumer
	}
	tests := []struct {
		name string
		ls   *localStorage
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ls.RegisterCacheConsumer(tt.args.consumer)
		})
	}
}

func Test_localStorage_cacheMaintainWorker(t *testing.T) {
	tests := []struct {
		name string
		ls   *localStorage
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ls.cacheMaintainWorker()
		})
	}
}

func Test_localStorage_readFromCache(t *testing.T) {
	type args struct {
		columnName string
		tableName  string
		date       time.Time
	}
	tests := []struct {
		name string
		ls   *localStorage
		args args
		want interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ls.readFromCache(tt.args.columnName, tt.args.tableName, tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("localStorage.readFromCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_localStorage_ReadValue(t *testing.T) {
	type args struct {
		columnName string
		tableName  string
		date       time.Time
	}
	tests := []struct {
		name string
		ls   *localStorage
		args args
		want interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ls.ReadValue(tt.args.columnName, tt.args.tableName, tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("localStorage.ReadValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_localStorage_ReadValueNoCache(t *testing.T) {
	type args struct {
		columnName string
		tableName  string
		date       time.Time
	}
	tests := []struct {
		name string
		ls   *localStorage
		args args
		want interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ls.ReadValueNoCache(tt.args.columnName, tt.args.tableName, tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("localStorage.ReadValueNoCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
