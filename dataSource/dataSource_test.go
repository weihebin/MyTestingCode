package dataSource

import "testing"

func Test_dataProvider_RegisterColumn(t *testing.T) {
	dp := NewDataProvider()
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

type testColumn struct {
	name string
}

func (c *testColumn) GetRequiredColumn() []IColumn {
	return nil

}
func (c *testColumn) GetName() string {
	return c.name
}
func BenchmarkRegisterColumn(b *testing.B) {
	dp := NewDataProvider()
	tc := &testColumn{name: "ABC"}
	dp.RegisterColumn(tc)
}
