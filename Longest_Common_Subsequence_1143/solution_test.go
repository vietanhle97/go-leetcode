package Longest_Common_Subsequence_1143

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LCS(t *testing.T) {

	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "LCS",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "LCS",
			args: args{
				text1: "abcdfgasgsadfgasge",
				text2: "asdfgsdfgsdfgsdfgsdce",
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestCommonSubsequence(tt.args.text1, tt.args.text2))
		})
	}
}
