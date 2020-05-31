package Asteroid_Collision_735

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "asteroid collision",
			args: args{
				asteroids: []int{5, 10, -5},
			},
			want: []int{5, 10},
		},
		{
			name: "asteroid collision",
			args: args{
				asteroids: []int{8, -8},
			},
			want: []int{},
		},
		{
			name: "asteroid collision",
			args: args{
				asteroids: []int{10, 2, -5},
			},
			want: []int{10},
		},
		{
			name: "asteroid collision",
			args: args{
				asteroids: []int{-2, -1, 1, 2},
			},
			want: []int{-2, -1, 1, 2},
		},
		{
			name: "asteroid collision",
			args: args{
				asteroids: []int{-2, -2, -2, 1},
			},
			want: []int{-2, -2, -2, 1},
		},
		{
			name: "asteroid collision",
			args: args{
				asteroids: []int{-2, -2, -2, -2},
			},
			want: []int{-2, -2, -2, -2},
		},
		{
			name: "asteroid collision",
			args: args{
				asteroids: []int{-2, -1, 1, 2},
			},
			want: []int{-2, -1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, asteroidCollision(tt.args.asteroids))
		})
	}
}
