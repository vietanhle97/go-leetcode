package Palindrome_Partitioning_131

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_palindromePartition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "palindrome partition I",
			args: args{
				s: "ababbbabbaba",
			},
			want: [][]string{{"a", "b", "a", "b", "b", "b", "a", "b", "b", "a", "b", "a"}, {"a", "b", "a", "b", "b", "b", "a", "b", "b", "aba"}, {"a", "b", "a", "b", "b", "b", "a", "b", "bab", "a"}, {"a", "b", "a", "b", "b", "b", "a", "bb", "a", "b", "a"}, {"a", "b", "a", "b", "b", "b", "a", "bb", "aba"}, {"a", "b", "a", "b", "b", "b", "abba", "b", "a"}, {"a", "b", "a", "b", "b", "bab", "b", "a", "b", "a"}, {"a", "b", "a", "b", "b", "bab", "b", "aba"}, {"a", "b", "a", "b", "b", "bab", "bab", "a"}, {"a", "b", "a", "b", "b", "babbab", "a"}, {"a", "b", "a", "b", "bb", "a", "b", "b", "a", "b", "a"}, {"a", "b", "a", "b", "bb", "a", "b", "b", "aba"}, {"a", "b", "a", "b", "bb", "a", "b", "bab", "a"}, {"a", "b", "a", "b", "bb", "a", "bb", "a", "b", "a"}, {"a", "b", "a", "b", "bb", "a", "bb", "aba"}, {"a", "b", "a", "b", "bb", "abba", "b", "a"}, {"a", "b", "a", "b", "bbabb", "a", "b", "a"}, {"a", "b", "a", "b", "bbabb", "aba"}, {"a", "b", "a", "bb", "b", "a", "b", "b", "a", "b", "a"}, {"a", "b", "a", "bb", "b", "a", "b", "b", "aba"}, {"a", "b", "a", "bb", "b", "a", "b", "bab", "a"}, {"a", "b", "a", "bb", "b", "a", "bb", "a", "b", "a"}, {"a", "b", "a", "bb", "b", "a", "bb", "aba"}, {"a", "b", "a", "bb", "b", "abba", "b", "a"}, {"a", "b", "a", "bb", "bab", "b", "a", "b", "a"}, {"a", "b", "a", "bb", "bab", "b", "aba"}, {"a", "b", "a", "bb", "bab", "bab", "a"}, {"a", "b", "a", "bb", "babbab", "a"}, {"a", "b", "a", "bbb", "a", "b", "b", "a", "b", "a"}, {"a", "b", "a", "bbb", "a", "b", "b", "aba"}, {"a", "b", "a", "bbb", "a", "b", "bab", "a"}, {"a", "b", "a", "bbb", "a", "bb", "a", "b", "a"}, {"a", "b", "a", "bbb", "a", "bb", "aba"}, {"a", "b", "a", "bbb", "abba", "b", "a"}, {"a", "b", "abbba", "b", "b", "a", "b", "a"}, {"a", "b", "abbba", "b", "b", "aba"}, {"a", "b", "abbba", "b", "bab", "a"}, {"a", "b", "abbba", "bb", "a", "b", "a"}, {"a", "b", "abbba", "bb", "aba"}, {"a", "bab", "b", "b", "a", "b", "b", "a", "b", "a"}, {"a", "bab", "b", "b", "a", "b", "b", "aba"}, {"a", "bab", "b", "b", "a", "b", "bab", "a"}, {"a", "bab", "b", "b", "a", "bb", "a", "b", "a"}, {"a", "bab", "b", "b", "a", "bb", "aba"}, {"a", "bab", "b", "b", "abba", "b", "a"}, {"a", "bab", "b", "bab", "b", "a", "b", "a"}, {"a", "bab", "b", "bab", "b", "aba"}, {"a", "bab", "b", "bab", "bab", "a"}, {"a", "bab", "b", "babbab", "a"}, {"a", "bab", "bb", "a", "b", "b", "a", "b", "a"}, {"a", "bab", "bb", "a", "b", "b", "aba"}, {"a", "bab", "bb", "a", "b", "bab", "a"}, {"a", "bab", "bb", "a", "bb", "a", "b", "a"}, {"a", "bab", "bb", "a", "bb", "aba"}, {"a", "bab", "bb", "abba", "b", "a"}, {"a", "bab", "bbabb", "a", "b", "a"}, {"a", "bab", "bbabb", "aba"}, {"a", "babbbab", "b", "a", "b", "a"}, {"a", "babbbab", "b", "aba"}, {"a", "babbbab", "bab", "a"}, {"aba", "b", "b", "b", "a", "b", "b", "a", "b", "a"}, {"aba", "b", "b", "b", "a", "b", "b", "aba"}, {"aba", "b", "b", "b", "a", "b", "bab", "a"}, {"aba", "b", "b", "b", "a", "bb", "a", "b", "a"}, {"aba", "b", "b", "b", "a", "bb", "aba"}, {"aba", "b", "b", "b", "abba", "b", "a"}, {"aba", "b", "b", "bab", "b", "a", "b", "a"}, {"aba", "b", "b", "bab", "b", "aba"}, {"aba", "b", "b", "bab", "bab", "a"}, {"aba", "b", "b", "babbab", "a"}, {"aba", "b", "bb", "a", "b", "b", "a", "b", "a"}, {"aba", "b", "bb", "a", "b", "b", "aba"}, {"aba", "b", "bb", "a", "b", "bab", "a"}, {"aba", "b", "bb", "a", "bb", "a", "b", "a"}, {"aba", "b", "bb", "a", "bb", "aba"}, {"aba", "b", "bb", "abba", "b", "a"}, {"aba", "b", "bbabb", "a", "b", "a"}, {"aba", "b", "bbabb", "aba"}, {"aba", "bb", "b", "a", "b", "b", "a", "b", "a"}, {"aba", "bb", "b", "a", "b", "b", "aba"}, {"aba", "bb", "b", "a", "b", "bab", "a"}, {"aba", "bb", "b", "a", "bb", "a", "b", "a"}, {"aba", "bb", "b", "a", "bb", "aba"}, {"aba", "bb", "b", "abba", "b", "a"}, {"aba", "bb", "bab", "b", "a", "b", "a"}, {"aba", "bb", "bab", "b", "aba"}, {"aba", "bb", "bab", "bab", "a"}, {"aba", "bb", "babbab", "a"}, {"aba", "bbb", "a", "b", "b", "a", "b", "a"}, {"aba", "bbb", "a", "b", "b", "aba"}, {"aba", "bbb", "a", "b", "bab", "a"}, {"aba", "bbb", "a", "bb", "a", "b", "a"}, {"aba", "bbb", "a", "bb", "aba"}, {"aba", "bbb", "abba", "b", "a"}},
		},
		{
			name: "palindrome partition I",
			args: args{
				s: "aab",
			},
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name: "palindrome partition I",
			args: args{
				s: "aaabbaaa",
			},
			want: [][]string{{"a", "a", "a", "b", "b", "a", "a", "a"}, {"a", "a", "a", "b", "b", "a", "aa"}, {"a", "a", "a", "b", "b", "aa", "a"}, {"a", "a", "a", "b", "b", "aaa"}, {"a", "a", "a", "bb", "a", "a", "a"}, {"a", "a", "a", "bb", "a", "aa"}, {"a", "a", "a", "bb", "aa", "a"}, {"a", "a", "a", "bb", "aaa"}, {"a", "a", "abba", "a", "a"}, {"a", "a", "abba", "aa"}, {"a", "aa", "b", "b", "a", "a", "a"}, {"a", "aa", "b", "b", "a", "aa"}, {"a", "aa", "b", "b", "aa", "a"}, {"a", "aa", "b", "b", "aaa"}, {"a", "aa", "bb", "a", "a", "a"}, {"a", "aa", "bb", "a", "aa"}, {"a", "aa", "bb", "aa", "a"}, {"a", "aa", "bb", "aaa"}, {"a", "aabbaa", "a"}, {"aa", "a", "b", "b", "a", "a", "a"}, {"aa", "a", "b", "b", "a", "aa"}, {"aa", "a", "b", "b", "aa", "a"}, {"aa", "a", "b", "b", "aaa"}, {"aa", "a", "bb", "a", "a", "a"}, {"aa", "a", "bb", "a", "aa"}, {"aa", "a", "bb", "aa", "a"}, {"aa", "a", "bb", "aaa"}, {"aa", "abba", "a", "a"}, {"aa", "abba", "aa"}, {"aaa", "b", "b", "a", "a", "a"}, {"aaa", "b", "b", "a", "aa"}, {"aaa", "b", "b", "aa", "a"}, {"aaa", "b", "b", "aaa"}, {"aaa", "bb", "a", "a", "a"}, {"aaa", "bb", "a", "aa"}, {"aaa", "bb", "aa", "a"}, {"aaa", "bb", "aaa"}, {"aaabbaaa"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, partition(tt.args.s))
		})
	}
}
