package multiply_strings

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ex1",
			args: args{
				num1: "2",
				num2: "3",
			},
			want: "6",
		},
		{
			name: "ex2",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "56088",
		},
		{
			name: "ex3",
			args: args{
				num1: "456",
				num2: "789",
			},
			want: "359784",
		},
		{
			name: "ex4",
			args: args{
				num1: "4561231241241241241",
				num2: "7891231241241241240",
			},
			want: "35993730469408447579160586257117978840",
		},
		{
			name: "ex5",
			args: args{
				num1: "4561231241241241241",
				num2: "0",
			},
			want: "0",
		},
		{
			name: "ex6",
			args: args{
				num1: "0",
				num2: "1230451316399123",
			},
			want: "0",
		},
		{
			name: "ex7",
			args: args{
				num1: "0",
				num2: "0",
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
