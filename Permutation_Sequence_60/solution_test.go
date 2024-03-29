package Permutation_Sequence_60

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "combination 3",
			args: args{
				n: 6,
				k: 80,
			},
			want: "153264",
		},
		{
			name: "combination 3",
			args: args{
				n: 9,
				k: 3000,
			},
			want: "127398654",
		},
		{
			name: "combination 3",
			args: args{
				n: 8,
				k: 600,
			},
			want: "12786543",
		},
		{
			name: "combination 3",
			args: args{
				n: 9,
				k: 241920,
			},
			want: "698754321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getPermutation(tt.args.n, tt.args.k))
		})
	}
}
