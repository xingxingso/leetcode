package longest_word_in_dictionary_through_deleting

import "testing"

func Test_findLongestWord(t *testing.T) {
	type args struct {
		s          string
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ex1",
			args: args{
				s:          "abpcplea",
				dictionary: []string{"ale", "apple", "monkey", "plea"},
			},
			want: "apple",
		},
		{
			name: "ex2",
			args: args{
				s:          "abpcplea",
				dictionary: []string{"a", "b", "c"},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestWord(tt.args.s, tt.args.dictionary); got != tt.want {
				t.Errorf("findLongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
