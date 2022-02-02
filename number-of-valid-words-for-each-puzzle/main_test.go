package number_of_valid_words_for_each_puzzle

import (
	"reflect"
	"testing"
)

func Test_findNumOfValidWords(t *testing.T) {
	type args struct {
		words   []string
		puzzles []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ex1",
			args: args{
				words:   []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},
				puzzles: []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"},
			},
			want: []int{1, 1, 3, 2, 4, 0},
		},
		{
			name: "ex2",
			args: args{
				words:   []string{"faced", "cabbage", "baggage", "beefed", "based"},
				puzzles: []string{"abcdefg"},
			},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumOfValidWords(tt.args.words, tt.args.puzzles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumOfValidWords() = %v, want %v", got, tt.want)
			}

			if got := findNumOfValidWordsO1(tt.args.words, tt.args.puzzles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumOfValidWords() = %v, want %v", got, tt.want)
			}

			if got := findNumOfValidWordsO2(tt.args.words, tt.args.puzzles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumOfValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
