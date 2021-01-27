package shortest_word_distance

import "testing"

func Test_shortestDistance(t *testing.T) {
	type args struct {
		words []string
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "equal",
			args: args{
				words: []string{"practice", "makes", "perfect", "coding", "makes"},
				word1: "coding",
				word2: "practice",
			},
			want: 3,
		},
		{
			name: "equal",
			args: args{
				words: []string{"practice", "makes", "perfect", "coding", "makes"},
				word1: "makes",
				word2: "coding",
			},
			want: 1,
		},
		{
			name: "equal",
			args: args{
				words: []string{"practice", "makes"},
				word1: "makes",
				word2: "practice",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDistance(tt.args.words, tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("shortestDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
