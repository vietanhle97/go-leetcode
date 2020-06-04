package Next_Greater_Element_III_556

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "combination 3",
			args: args{
				n: 11,
			},
			want: -1,
		},
		{
			name: "combination 3",
			args: args{
				n: 12,
			},
			want: 21,
		},
		{
			name: "combination 3",
			args: args{
				n: 16,
			},
			want: 61,
		},
		{
			name: "combination 3",
			args: args{
				n: 15,
			},
			want: 51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, nextGreaterElement(tt.args.n))
		})
	}
}
