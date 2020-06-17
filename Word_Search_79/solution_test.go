package Word_Search_79

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_wordSearch(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "word search",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "ABCCED",
			},
			want: true,
		},
		{
			name: "word search",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "SEE",
			},
			want: true,
		},
		{
			name: "word search",
			args: args{
				board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
				word:  "ABCD",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, exist(tt.args.board, tt.args.word))
		})
	}
}
