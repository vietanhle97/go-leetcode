package Rectangle_Area_223

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		A int
		B int
		C int
		D int
		E int
		F int
		G int
		H int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "coin change",
			args: args{
				A: -3,
				B: 0,
				C: 3,
				D: 4,
				E: 0,
				F: -1,
				G: 9,
				H: 2,
			},
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, computeArea(tt.args.A, tt.args.B, tt.args.C, tt.args.D, tt.args.E, tt.args.F, tt.args.G, tt.args.H))
		})
	}
}
