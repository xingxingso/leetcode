package excel_sheet_column_title

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		columnNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ex1",
			args: args{
				columnNumber: 1,
			},
			want: "A",
		},
		{
			name: "ex2",
			args: args{
				columnNumber: 28,
			},
			want: "AB",
		},
		{
			name: "ex3",
			args: args{
				columnNumber: 701,
			},
			want: "ZY",
		},
		{
			name: "ex4",
			args: args{
				columnNumber: 2147483647,
			},
			want: "FXSHRXW",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
