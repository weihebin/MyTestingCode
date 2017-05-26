package localStorage

import (
	"MyTestingCode/priceData"
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
	catalog   string
	name      string
	value     interface{}
	valueType priceData.DataValueType
}

func (tc *testColumn) GetName() string {
	return tc.name
}
func (tc *testColumn) GetCatalog() string {
	return tc.catalog
}
func (tc *testColumn) GetValueType() priceData.DataValueType {
	return tc.valueType
}
func getTestingColumn(name string) ILocalStorageDataInfo {
	switch name {
	case "Closed":
		return &testColumn{name: name, valueType: priceData.DataValueType_Decimal}
	}
	return nil
}
func TestHappyPath(t *testing.T) {
	ls := setupTestLocalStorage()

	ls.Register(getTestingColumn("Closed"))

	err := ls.LoadData("NVDA")
	defer ls.Close()
	assert.Nil(t, err)
	assert.NotNil(t, ls.database)

}
