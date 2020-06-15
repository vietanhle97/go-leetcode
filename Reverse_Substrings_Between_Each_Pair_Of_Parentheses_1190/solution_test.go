package Reverse_Substrings_Between_Each_Pair_Of_Parentheses_1190

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverse string in parenthesis",
			args: args{
				s: "(abcd)",
			},
			want: "dcba",
		},
		{
			name: "reverse string in parenthesis",
			args: args{
				s: "(u(love)i)",
			},
			want: "iloveu",
		},
		{
			name: "reverse string in parenthesis",
			args: args{
				s: "(ed(et(oc))el)",
			},
			want: "leetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseParentheses(tt.args.s))
		})
	}
}
