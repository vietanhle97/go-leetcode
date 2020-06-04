package Word_Break_II_140

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "word break",
			args: args{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			want: []string{"cat sand dog", "cats and dog"},
		},
		{
			name: "word break",
			args: args{
				s:        "pineapplepenapple",
				wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			},
			want: []string{"pine applepen apple", "pineapple pen apple", "pine apple pen apple"},
		},
		{
			name: "word break",
			args: args{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			want: []string{"leet code"},
		},
		{
			name: "word break",
			args: args{
				s:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, wordBreak(tt.args.s, tt.args.wordDict))
		})
	}
}
