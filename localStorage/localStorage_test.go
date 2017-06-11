package localStorage

import (
	"MyTestingCode/constants"
	"time"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_localStorage_getLocation(t *testing.T) {
	ls := setupTestLocalStorage()
	os.Setenv("LOCAL_STORAGE_ROOT", "root")
	type args struct {
		symbol string
	}
	tests := []struct {
		name string
		ls   *sqliteLocalStorage
		args args
		want string
	}{
		{
			name: "GetLocation",
			ls:   ls,
			args: args{symbol: "abcd"},
			want: "root/a/abcd/abcd.db",
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
func setupTestLocalStorage() *sqliteLocalStorage {
	os.Setenv("LOCAL_STORAGE_ROOT", "C:/Users/hwei/AppData/Local/Temp/")

	return &sqliteLocalStorage{}

}

type testColumn struct {
	catalog   constants.ColumnCatalog
	name      string
	value     interface{}
	valueType constants.DataValueType
}

func (tc *testColumn) GetName() string {
	return tc.name
}
func (tc *testColumn) GetCatalog() constants.ColumnCatalog {
	return tc.catalog
}
func (tc *testColumn) GetValueType() constants.DataValueType {
	return tc.valueType
}
func getTestingColumn(name string) *testColumn {
	switch name {
	// case "Closed":
	// 	return &testColumn{name: name, valueType: constants.DataValueType_Decimal, catalog: constants.ColumnCatalog_Basic}
	// case "Open":
	// 	return &testColumn{name: name, valueType: constants.DataValueType_Decimal, catalog: constants.ColumnCatalog_Basic}
	default:
		return &testColumn{name: name, valueType: constants.DataValueType_Decimal, catalog: constants.ColumnCatalog_Basic}

	}
	return nil
}
func TestHappyPath(t *testing.T) {
	ls := setupTestLocalStorage()
	column := getTestingColumn("Closed")
	ls.RegisterColumn(column.GetCatalog(), column.GetName(), column.GetValueType())
	column = getTestingColumn("Open")
	ls.RegisterColumn(column.GetCatalog(), column.GetName(), column.GetValueType())
	column = getTestingColumn("High")
	ls.RegisterColumn(column.GetCatalog(), column.GetName(), column.GetValueType())
	err := ls.LoadData("NVDA")
	defer ls.Close()
	assert.Nil(t, err)
	assert.NotNil(t, ls.database)

	sD := ls.GetLastDateOfColumn(column.GetCatalog(), column.GetName())
	// no data
	assert.Equal(t, 1, sD.Year())
	assert.Equal(t, time.Month(1), sD.Month())

	assert.Equal(t, 1, sD.Day())

}

func BenchmarkGetLastDateOfColumn(b *testing.B) {

	ls := setupTestLocalStorage()
	column := getTestingColumn("Closed")
	ls.RegisterColumn(column.GetCatalog(), column.GetName(), column.GetValueType())

	ls.LoadData("NVDA")

	defer ls.Close()

	for n := 0; n <= b.N; n++ {
		ls.GetLastDateOfColumn(column.GetCatalog(), column.GetName())

	}
}
