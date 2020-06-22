package Longest_Substring_With_At_Least_K_Repeating_Characters_395

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestSubStringWithDupChar(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "longest substring with at least K repeating characters",
			args: args{
				s: "zzzzzzzzzzaaaaaaaaabbbbbbbbhbhbhbhbhbhbhicbcbcibcbccccccccccbbbbbbbbaaaaaaaaafffaahhhhhiaahiiiiiiiiifeeeeeeeeee",
				k: 10,
			},
			want: 21,
		},
		{
			name: "longest substring with at least K repeating characters",
			args: args{
				s: "aaabbbcdefcdefgggggggggggggggcde",
				k: 3,
			},
			want: 15,
		},
		{
			name: "longest substring with at least K repeating characters",
			args: args{
				s: "aaabbbcdefcdefcde",
				k: 3,
			},
			want: 6,
		},
		{
			name: "longest substring with at least K repeating characters",
			args: args{
				s: "ababbc",
				k: 2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestSubstring(tt.args.s, tt.args.k))
		})
	}
}
