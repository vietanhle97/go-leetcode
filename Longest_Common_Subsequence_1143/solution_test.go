package Longest_Common_Subsequence_1143

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LCS(t *testing.T) {

	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "LCS",
			args: args{
				s1: "abcde",
				s2: "ace",
			},
			want: 3,
		},
		{
			name: "LCS",
			args: args{
				s1: "abcdfgasgsadfgasge",
				s2: "asdfgsdfgsdfgsdfgsdce",
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestCommonSubsequence(tt.args.s1, tt.args.s2))
		})
	}
}
