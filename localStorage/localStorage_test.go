package localStorage

import (
	"os"
	"testing"
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

type testColumn struct {
	name string
}

func getTestingColumn(name string) ICacheDataConsumer {
	switch name {
	case "Closed":
		return &testColumn{name: name}
	}
	return nil
}
func TestHappyPath(t *testing.T) {
	ls := setupTestLocalStorage()

	ls.RegisterCacheConsumer(getTestingColumn("Closed"))

}
