package Maximum_Side_Length_Of_A_Square_With_Sum_Less_Than_Or_Equal_To_Threshold_1292

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxLengthWithSumLessThanThreshold(t *testing.T) {
	type args struct {
		mat       [][]int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "maximum side length of a square with sum less than or equal to threshold",
			args: args{
				mat:       [][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}},
				threshold: 4,
			},
			want: 2,
		},
		{
			name: "maximum side length of a square with sum less than or equal to threshold",
			args: args{
				mat:       [][]int{{28, 39, 98, 91, 7, 99}, {79, 3, 17, 83, 9, 92}, {81, 73, 42, 27, 67, 70}, {88, 30, 73, 99, 96, 89}, {27, 59, 0, 1, 65, 79}, {42, 55, 48, 29, 86, 96}},
				threshold: 24829,
			},
			want: 6,
		},
		{
			name: "maximum side length of a square with sum less than or equal to threshold",
			args: args{
				mat:       [][]int{{75, 5, 88, 23, 94, 89, 14, 43, 34, 22, 8, 70, 51, 64, 7, 86, 71, 54, 66, 53, 73, 99, 38, 68, 59, 18, 24, 2, 28, 57, 42, 10, 25, 35, 60, 11, 65, 85, 40, 47, 72, 91, 97, 13, 41, 21, 82, 76}, {79, 87, 27, 9, 68, 86, 10, 99, 97, 36, 43, 7, 100, 54, 95, 48, 41, 34, 70, 75, 58, 12, 82, 96, 65, 40, 8, 16, 24, 61, 38, 39, 67, 31, 49, 98, 78, 46, 6, 71, 72, 94, 51, 76, 32, 0, 20, 42}, {15, 11, 64, 98, 22, 49, 6, 89, 70, 27, 12, 88, 72, 42, 66, 95, 57, 68, 2, 79, 53, 10, 9, 97, 60, 78, 76, 14, 61, 30, 85, 50, 37, 56, 41, 26, 59, 92, 18, 46, 55, 86, 80, 67, 29, 54, 17, 81}, {6, 53, 97, 30, 99, 2, 73, 24, 59, 9, 95, 100, 70, 90, 79, 19, 37, 64, 47, 68, 38, 98, 44, 11, 0, 28, 43, 54, 81, 5, 62, 67, 25, 29, 80, 14, 45, 57, 42, 93, 78, 83, 76, 22, 7, 33, 39, 21}, {81, 15, 20, 62, 83, 26, 54, 51, 56, 78, 43, 87, 44, 76, 58, 92, 36, 6, 55, 100, 4, 41, 24, 94, 93, 8, 10, 0, 46, 5, 61, 99, 89, 11, 59, 45, 16, 34, 50, 90, 66, 98, 84, 3, 12, 86, 37, 17}, {57, 85, 51, 50, 5, 70, 25, 0, 17, 59, 91, 52, 48, 32, 37, 4, 72, 53, 13, 33, 34, 27, 98, 80, 67, 9, 90, 73, 86, 44, 92, 42, 6, 81, 97, 21, 89, 82, 69, 20, 76, 96, 8, 49, 100, 83, 63, 79}, {66, 18, 33, 48, 44, 84, 25, 91, 3, 71, 21, 23, 64, 15, 88, 75, 62, 99, 6, 34, 50, 51, 41, 49, 2, 19, 11, 22, 5, 57, 28, 95, 56, 29, 35, 53, 4, 55, 43, 13, 92, 87, 81, 89, 27, 14, 17, 30}, {54, 75, 28, 38, 94, 10, 45, 73, 21, 66, 74, 9, 63, 33, 40, 83, 44, 37, 91, 72, 65, 69, 49, 70, 1, 92, 16, 87, 80, 8, 82, 96, 18, 100, 62, 61, 52, 12, 50, 17, 59, 19, 29, 20, 25, 95, 15, 86}, {60, 54, 50, 20, 36, 85, 47, 97, 16, 30, 76, 96, 100, 26, 81, 4, 95, 0, 48, 21, 87, 19, 45, 34, 92, 1, 33, 74, 8, 41, 5, 90, 37, 70, 29, 23, 72, 64, 78, 89, 27, 13, 22, 18, 80, 61, 56, 66}, {5, 75, 1, 76, 94, 95, 24, 100, 81, 35, 84, 28, 86, 18, 55, 54, 57, 51, 4, 97, 2, 42, 41, 83, 96, 58, 59, 9, 8, 39, 92, 61, 69, 13, 14, 44, 32, 60, 79, 25, 45, 62, 21, 20, 98, 53, 68, 0}, {91, 14, 90, 50, 33, 29, 57, 63, 73, 46, 36, 71, 66, 19, 74, 27, 84, 22, 28, 52, 54, 30, 25, 75, 64, 4, 100, 44, 58, 0, 9, 31, 88, 72, 32, 95, 13, 24, 43, 17, 70, 5, 10, 18, 16, 7, 3, 80}, {88, 69, 57, 28, 27, 82, 61, 18, 25, 0, 64, 13, 29, 74, 35, 90, 1, 92, 46, 16, 32, 41, 40, 23, 49, 70, 48, 14, 94, 91, 78, 17, 60, 87, 37, 96, 9, 6, 55, 72, 95, 12, 99, 62, 98, 51, 43, 89}, {26, 33, 15, 99, 97, 10, 87, 98, 93, 48, 72, 51, 12, 63, 81, 75, 24, 29, 69, 57, 43, 42, 40, 88, 1, 76, 55, 39, 95, 70, 3, 52, 85, 56, 82, 28, 62, 17, 18, 89, 20, 59, 22, 84, 14, 79, 46, 66}, {37, 57, 30, 84, 44, 3, 90, 18, 82, 91, 6, 9, 94, 66, 80, 36, 63, 13, 74, 40, 59, 92, 29, 15, 45, 77, 70, 27, 19, 71, 10, 46, 68, 12, 42, 93, 47, 58, 33, 88, 87, 98, 97, 72, 89, 81, 75, 7}, {31, 70, 0, 73, 93, 55, 88, 92, 33, 68, 65, 39, 49, 30, 27, 44, 71, 52, 12, 78, 75, 47, 46, 69, 15, 18, 74, 38, 5, 87, 77, 37, 21, 89, 13, 62, 80, 58, 66, 20, 41, 3, 45, 97, 9, 11, 50, 100}, {59, 78, 87, 27, 82, 1, 10, 60, 47, 98, 17, 19, 89, 23, 96, 53, 95, 40, 93, 52, 8, 13, 9, 88, 30, 62, 100, 61, 68, 16, 86, 71, 29, 91, 76, 14, 80, 28, 50, 57, 7, 74, 15, 26, 20, 46, 81, 42}, {88, 59, 52, 48, 38, 7, 66, 47, 46, 5, 68, 63, 2, 42, 22, 85, 49, 80, 18, 6, 61, 76, 82, 93, 65, 19, 86, 87, 12, 15, 10, 34, 40, 70, 54, 36, 75, 11, 43, 83, 3, 45, 81, 96, 64, 71, 16, 100}, {48, 99, 27, 84, 47, 6, 33, 8, 11, 65, 95, 16, 88, 68, 23, 17, 19, 9, 49, 89, 2, 13, 59, 77, 81, 31, 69, 20, 7, 35, 56, 15, 44, 97, 32, 39, 62, 43, 30, 96, 46, 21, 58, 71, 61, 54, 28, 3}, {97, 86, 8, 98, 59, 37, 51, 53, 9, 0, 36, 19, 55, 24, 44, 3, 64, 22, 15, 93, 63, 29, 99, 100, 52, 13, 89, 49, 16, 82, 11, 73, 54, 34, 80, 70, 12, 28, 7, 90, 26, 71, 35, 78, 85, 41, 88, 65}, {57, 100, 84, 81, 53, 46, 8, 61, 91, 27, 69, 50, 28, 76, 48, 89, 37, 94, 47, 19, 10, 21, 92, 80, 87, 26, 15, 24, 63, 85, 35, 38, 43, 42, 55, 39, 58, 14, 99, 2, 13, 32, 52, 68, 20, 97, 70, 93}, {31, 98, 56, 51, 44, 23, 73, 41, 42, 83, 96, 82, 49, 92, 88, 72, 4, 63, 74, 64, 86, 69, 13, 59, 80, 5, 19, 57, 61, 8, 97, 9, 70, 85, 16, 78, 1, 32, 3, 53, 71, 33, 48, 2, 34, 6, 20, 58}, {19, 91, 70, 6, 38, 15, 54, 90, 17, 10, 56, 8, 93, 74, 53, 35, 85, 73, 79, 60, 63, 24, 23, 84, 87, 48, 89, 68, 34, 42, 13, 39, 82, 27, 76, 69, 98, 2, 16, 11, 92, 77, 83, 75, 46, 65, 86, 100}, {38, 22, 40, 97, 75, 66, 39, 69, 90, 89, 47, 25, 99, 45, 65, 9, 76, 15, 53, 96, 1, 55, 28, 57, 77, 35, 70, 0, 84, 73, 93, 48, 50, 100, 42, 64, 19, 79, 67, 92, 36, 11, 80, 86, 29, 82, 61, 12}, {8, 6, 18, 34, 61, 95, 68, 70, 94, 48, 10, 90, 57, 56, 37, 20, 45, 79, 96, 53, 76, 88, 99, 78, 14, 67, 13, 85, 74, 64, 55, 21, 71, 69, 98, 27, 86, 49, 72, 60, 1, 26, 23, 80, 97, 59, 7, 63}, {67, 18, 26, 47, 13, 3, 12, 58, 97, 10, 57, 84, 23, 29, 43, 62, 94, 44, 36, 7, 22, 100, 64, 11, 66, 0, 56, 83, 25, 9, 90, 20, 95, 6, 59, 33, 15, 87, 52, 41, 8, 86, 88, 5, 42, 82, 46, 50}, {0, 93, 49, 66, 69, 1, 39, 76, 77, 73, 75, 6, 32, 84, 29, 3, 61, 5, 34, 56, 100, 85, 18, 86, 40, 90, 9, 4, 74, 11, 37, 45, 44, 25, 98, 72, 30, 99, 27, 33, 83, 35, 89, 17, 15, 62, 65, 14}, {74, 68, 53, 54, 65, 33, 37, 44, 55, 78, 22, 49, 71, 70, 15, 93, 16, 66, 10, 62, 12, 52, 7, 40, 1, 9, 18, 94, 80, 83, 47, 28, 14, 46, 19, 24, 100, 34, 81, 95, 79, 59, 75, 2, 5, 29, 39, 99}, {31, 34, 18, 3, 22, 55, 63, 72, 71, 2, 23, 86, 89, 44, 32, 39, 24, 12, 60, 81, 95, 92, 46, 28, 37, 82, 26, 83, 40, 74, 17, 94, 99, 59, 13, 56, 16, 93, 85, 5, 0, 75, 49, 8, 38, 57, 4, 80}, {75, 71, 27, 79, 21, 37, 91, 74, 90, 6, 9, 67, 98, 66, 15, 56, 99, 7, 68, 5, 3, 96, 81, 57, 46, 45, 23, 58, 78, 18, 47, 93, 28, 2, 34, 83, 25, 32, 72, 13, 80, 59, 16, 11, 48, 73, 39, 61}, {93, 45, 21, 54, 44, 51, 9, 10, 40, 5, 2, 98, 100, 20, 29, 62, 99, 42, 35, 73, 84, 53, 55, 85, 43, 49, 50, 37, 17, 16, 48, 32, 6, 11, 65, 74, 26, 7, 14, 66, 33, 89, 18, 25, 38, 30, 63, 4}, {87, 91, 43, 34, 85, 60, 50, 27, 29, 76, 61, 44, 15, 22, 63, 51, 46, 47, 88, 80, 14, 19, 8, 52, 41, 38, 20, 86, 68, 3, 84, 94, 71, 6, 92, 64, 65, 67, 66, 25, 18, 32, 58, 69, 97, 75, 83, 45}, {58, 12, 3, 69, 100, 54, 90, 9, 99, 52, 57, 10, 83, 92, 15, 67, 14, 97, 96, 45, 38, 95, 59, 82, 25, 35, 39, 5, 86, 19, 41, 72, 43, 13, 47, 48, 42, 26, 98, 7, 50, 76, 55, 77, 93, 16, 84, 23}, {76, 64, 56, 72, 1, 47, 43, 94, 63, 38, 10, 0, 51, 30, 65, 3, 20, 67, 21, 85, 78, 97, 59, 74, 41, 42, 14, 28, 69, 66, 33, 52, 6, 2, 44, 93, 40, 32, 61, 34, 91, 35, 7, 9, 57, 31, 96, 50}, {66, 35, 1, 62, 49, 33, 56, 96, 80, 97, 8, 76, 36, 75, 15, 18, 68, 63, 85, 67, 94, 37, 22, 60, 99, 87, 31, 52, 43, 64, 2, 98, 27, 58, 44, 17, 42, 86, 53, 4, 14, 46, 51, 12, 83, 79, 9, 30}, {47, 45, 8, 12, 3, 88, 31, 61, 28, 55, 10, 95, 93, 99, 5, 11, 13, 0, 59, 75, 54, 70, 24, 82, 90, 52, 36, 20, 37, 71, 100, 15, 87, 38, 67, 73, 21, 6, 9, 81, 19, 65, 97, 1, 94, 2, 62, 86}, {26, 29, 39, 69, 63, 66, 24, 91, 33, 59, 78, 81, 47, 70, 37, 94, 72, 4, 9, 40, 42, 34, 32, 20, 45, 88, 10, 15, 56, 48, 93, 6, 79, 5, 76, 36, 0, 50, 11, 52, 44, 58, 49, 86, 13, 35, 83, 60}, {37, 87, 95, 79, 3, 8, 80, 53, 52, 58, 75, 92, 39, 69, 88, 78, 98, 90, 10, 2, 4, 96, 62, 17, 72, 19, 7, 26, 67, 31, 5, 30, 48, 85, 34, 51, 0, 61, 11, 97, 16, 68, 89, 81, 65, 93, 20, 25}, {87, 99, 4, 5, 94, 13, 85, 76, 30, 33, 93, 80, 55, 53, 9, 75, 96, 100, 62, 46, 56, 25, 2, 20, 69, 61, 23, 89, 24, 43, 70, 59, 74, 97, 60, 18, 1, 36, 82, 91, 64, 40, 84, 15, 32, 68, 17, 35}, {16, 82, 9, 0, 37, 26, 88, 45, 21, 58, 64, 18, 43, 78, 68, 28, 52, 34, 94, 79, 15, 54, 23, 2, 32, 38, 62, 6, 3, 81, 100, 72, 92, 75, 5, 77, 33, 61, 14, 44, 55, 47, 99, 30, 4, 65, 46, 13}, {1, 96, 90, 63, 86, 30, 43, 3, 16, 4, 15, 8, 78, 71, 66, 57, 40, 36, 42, 51, 23, 88, 28, 17, 13, 87, 72, 83, 18, 94, 97, 33, 98, 52, 61, 79, 80, 11, 29, 39, 38, 31, 62, 81, 2, 64, 10, 12}, {21, 6, 77, 53, 17, 41, 25, 52, 5, 82, 70, 38, 27, 58, 54, 44, 28, 80, 97, 59, 68, 10, 89, 55, 61, 30, 18, 95, 9, 19, 65, 76, 72, 13, 36, 14, 15, 46, 60, 62, 35, 49, 85, 78, 48, 96, 84, 86}, {90, 70, 20, 29, 12, 41, 17, 87, 24, 57, 32, 34, 45, 75, 2, 35, 98, 11, 7, 73, 58, 66, 76, 16, 77, 38, 15, 55, 64, 85, 6, 62, 23, 40, 59, 71, 48, 51, 36, 63, 78, 99, 30, 42, 3, 39, 0, 95}, {16, 51, 99, 11, 89, 4, 33, 90, 29, 86, 82, 8, 17, 15, 66, 63, 68, 12, 84, 31, 93, 80, 26, 79, 10, 37, 20, 44, 71, 6, 2, 1, 3, 70, 23, 45, 50, 81, 91, 47, 75, 97, 28, 64, 24, 41, 92, 48}, {32, 66, 39, 43, 80, 14, 100, 37, 15, 61, 91, 94, 53, 62, 82, 95, 77, 65, 4, 30, 99, 84, 51, 26, 22, 44, 73, 56, 71, 78, 63, 17, 19, 85, 13, 23, 54, 47, 7, 10, 42, 16, 29, 57, 1, 21, 68, 88}, {29, 42, 55, 49, 9, 17, 24, 0, 32, 30, 100, 2, 89, 96, 40, 16, 53, 61, 6, 23, 73, 62, 98, 37, 58, 71, 94, 81, 84, 88, 45, 14, 95, 13, 19, 36, 52, 66, 59, 3, 28, 4, 92, 86, 64, 68, 20, 41}, {70, 86, 63, 2, 67, 42, 69, 62, 50, 18, 9, 66, 93, 8, 32, 30, 77, 87, 85, 41, 83, 46, 26, 80, 43, 7, 49, 95, 19, 58, 51, 52, 15, 96, 23, 99, 68, 61, 57, 90, 65, 84, 22, 5, 29, 35, 91, 27}, {69, 84, 71, 58, 76, 82, 63, 4, 61, 51, 2, 93, 37, 46, 53, 64, 19, 8, 34, 3, 23, 31, 60, 79, 27, 42, 41, 85, 96, 36, 78, 38, 91, 43, 22, 28, 68, 17, 100, 45, 83, 26, 44, 90, 16, 92, 95, 15}, {82, 16, 41, 64, 24, 44, 73, 48, 66, 75, 46, 43, 77, 97, 71, 1, 100, 0, 17, 40, 15, 89, 37, 20, 39, 79, 61, 67, 4, 91, 28, 96, 85, 35, 56, 23, 21, 19, 8, 87, 80, 52, 27, 31, 72, 83, 22, 94}, {5, 54, 82, 86, 100, 18, 62, 24, 11, 20, 84, 64, 61, 31, 48, 58, 47, 85, 76, 59, 77, 93, 42, 32, 51, 87, 67, 57, 8, 6, 75, 25, 72, 37, 30, 50, 27, 63, 56, 43, 70, 21, 71, 68, 22, 45, 96, 66}, {86, 16, 84, 74, 96, 32, 15, 25, 8, 30, 90, 48, 40, 100, 44, 91, 94, 64, 89, 6, 59, 83, 67, 79, 9, 52, 31, 53, 99, 92, 95, 60, 97, 13, 24, 68, 12, 3, 28, 0, 41, 81, 38, 18, 78, 50, 70, 43}, {82, 70, 45, 8, 79, 23, 39, 11, 81, 78, 92, 47, 19, 63, 57, 37, 35, 13, 55, 67, 6, 96, 65, 66, 1, 61, 29, 91, 97, 95, 88, 18, 2, 98, 16, 49, 32, 52, 48, 4, 80, 75, 94, 53, 56, 68, 3, 72}, {97, 99, 4, 6, 26, 10, 27, 28, 29, 38, 13, 40, 33, 43, 18, 86, 66, 75, 17, 41, 31, 34, 76, 3, 45, 67, 48, 11, 2, 88, 44, 59, 64, 94, 89, 91, 57, 49, 35, 62, 25, 16, 23, 56, 30, 61, 5, 82}, {42, 48, 17, 90, 61, 14, 54, 47, 70, 72, 24, 50, 93, 38, 2, 44, 57, 0, 78, 89, 23, 43, 66, 35, 76, 52, 34, 19, 32, 16, 45, 3, 92, 13, 91, 97, 99, 88, 33, 5, 11, 22, 58, 10, 8, 9, 83, 73}, {30, 91, 11, 9, 3, 21, 50, 19, 66, 59, 51, 75, 39, 60, 79, 0, 15, 85, 49, 41, 8, 37, 72, 54, 44, 96, 88, 14, 33, 31, 42, 35, 56, 74, 57, 17, 67, 87, 94, 20, 97, 73, 62, 26, 18, 25, 47, 13}, {64, 49, 2, 92, 81, 12, 43, 97, 44, 72, 42, 70, 67, 98, 41, 21, 10, 78, 61, 9, 76, 63, 26, 15, 73, 6, 39, 19, 40, 89, 87, 1, 14, 4, 79, 27, 7, 28, 56, 47, 85, 48, 77, 0, 8, 65, 94, 45}, {75, 30, 38, 77, 32, 84, 6, 53, 8, 57, 68, 96, 17, 0, 47, 78, 79, 40, 73, 52, 54, 95, 36, 13, 4, 43, 23, 71, 41, 87, 22, 89, 51, 86, 80, 42, 44, 14, 65, 94, 69, 99, 15, 49, 18, 31, 5, 67}, {64, 84, 31, 87, 23, 66, 99, 80, 38, 12, 5, 3, 24, 41, 63, 95, 79, 19, 33, 85, 16, 56, 54, 67, 9, 82, 26, 86, 25, 100, 28, 37, 68, 78, 0, 6, 21, 43, 57, 49, 83, 51, 8, 59, 13, 48, 2, 44}, {85, 79, 32, 1, 58, 2, 84, 19, 75, 37, 52, 30, 78, 25, 83, 20, 68, 40, 13, 97, 76, 29, 56, 12, 72, 16, 90, 69, 3, 61, 59, 77, 92, 44, 64, 91, 35, 48, 89, 63, 80, 74, 51, 82, 15, 26, 67, 10}, {73, 13, 76, 21, 98, 23, 58, 75, 25, 61, 70, 84, 59, 77, 86, 19, 44, 71, 100, 69, 24, 32, 46, 36, 8, 95, 26, 18, 40, 34, 81, 27, 9, 47, 83, 20, 94, 43, 1, 52, 30, 2, 33, 11, 28, 29, 87, 55}, {1, 70, 93, 44, 23, 90, 85, 28, 17, 39, 50, 29, 49, 72, 20, 35, 9, 14, 71, 13, 5, 78, 11, 52, 86, 26, 53, 73, 41, 43, 67, 19, 2, 68, 0, 56, 45, 84, 64, 83, 66, 97, 18, 32, 95, 8, 33, 54}, {24, 7, 23, 3, 93, 25, 32, 5, 78, 54, 80, 33, 36, 43, 39, 56, 8, 81, 50, 2, 97, 18, 61, 31, 52, 98, 9, 77, 90, 38, 45, 83, 41, 62, 51, 91, 15, 96, 40, 74, 95, 53, 76, 100, 22, 48, 12, 21}},
				threshold: 92003,
			},
			want: 43,
		},
		{
			name: "maximum side length of a square with sum less than or equal to threshold",
			args: args{
				mat:       [][]int{{21, 91, 65, 94, 46, 95, 70, 26, 88, 29, 73, 24, 62, 32, 35, 49, 79, 67, 23, 45, 6, 7, 60, 98, 31, 69, 86, 64, 28, 78, 47, 57, 20, 17}, {67, 74, 78, 10, 19, 22, 17, 58, 44, 31, 100, 37, 30, 16, 6, 56, 81, 60, 97, 41, 63, 87, 40, 79, 26, 77, 86, 80, 15, 39, 23, 28, 34, 83}, {24, 73, 55, 82, 0, 62, 89, 54, 49, 28, 77, 43, 81, 12, 84, 91, 72, 25, 5, 29, 68, 23, 93, 42, 34, 13, 69, 94, 35, 19, 37, 36, 78, 74}, {81, 27, 83, 11, 31, 0, 85, 71, 35, 4, 8, 92, 62, 99, 61, 47, 63, 17, 93, 7, 43, 52, 60, 33, 79, 41, 40, 76, 13, 80, 57, 39, 51, 54}, {21, 81, 28, 100, 47, 83, 5, 62, 60, 97, 86, 90, 85, 46, 0, 42, 72, 58, 74, 77, 38, 69, 41, 35, 95, 54, 8, 99, 1, 87, 33, 51, 66, 27}, {64, 69, 54, 34, 22, 21, 83, 82, 85, 2, 5, 81, 67, 56, 16, 4, 77, 84, 40, 72, 70, 95, 51, 89, 96, 61, 63, 7, 1, 28, 99, 62, 75, 47}, {2, 15, 95, 54, 89, 10, 96, 27, 6, 85, 1, 25, 81, 33, 37, 18, 39, 63, 49, 72, 77, 69, 48, 31, 24, 22, 87, 66, 8, 53, 17, 74, 86, 68}, {68, 88, 5, 53, 94, 65, 38, 95, 48, 37, 42, 55, 29, 93, 80, 47, 36, 77, 64, 46, 43, 2, 99, 89, 72, 61, 9, 16, 74, 20, 76, 73, 82, 11}, {28, 57, 76, 12, 86, 31, 37, 62, 70, 96, 27, 17, 13, 77, 59, 41, 85, 55, 94, 97, 3, 34, 56, 67, 74, 47, 33, 38, 6, 64, 51, 63, 83, 80}, {66, 41, 13, 1, 70, 81, 15, 0, 82, 92, 22, 12, 44, 74, 16, 34, 26, 20, 36, 35, 77, 53, 52, 79, 56, 29, 59, 10, 3, 89, 48, 80, 55, 24}, {92, 15, 98, 32, 23, 62, 68, 93, 1, 17, 53, 28, 21, 54, 60, 49, 51, 43, 52, 77, 3, 34, 46, 10, 84, 35, 75, 89, 26, 24, 80, 45, 31, 22}, {83, 95, 2, 51, 17, 7, 28, 89, 5, 80, 52, 43, 67, 91, 57, 86, 40, 53, 90, 11, 65, 38, 68, 47, 33, 82, 63, 85, 61, 37, 15, 22, 87, 19}, {2, 99, 69, 61, 0, 20, 50, 77, 48, 11, 22, 17, 28, 66, 1, 70, 98, 74, 88, 82, 63, 34, 75, 76, 78, 24, 23, 8, 10, 35, 26, 85, 19, 57}, {94, 18, 44, 87, 91, 81, 3, 15, 31, 40, 4, 14, 63, 47, 89, 49, 12, 16, 9, 23, 50, 66, 0, 28, 43, 10, 39, 20, 78, 56, 73, 99, 95, 5}, {94, 26, 89, 14, 69, 27, 8, 85, 87, 12, 24, 35, 44, 81, 15, 56, 38, 63, 75, 60, 99, 47, 32, 76, 10, 19, 62, 93, 20, 98, 70, 21, 3, 30}, {53, 35, 93, 3, 32, 84, 85, 68, 64, 77, 86, 59, 2, 65, 47, 12, 76, 5, 42, 61, 7, 37, 50, 0, 38, 79, 18, 52, 34, 82, 29, 55, 41, 21}, {96, 25, 50, 55, 69, 19, 15, 9, 30, 88, 16, 32, 38, 41, 80, 6, 97, 79, 28, 27, 65, 47, 89, 0, 73, 84, 3, 51, 11, 33, 1, 72, 23, 92}, {30, 65, 75, 51, 17, 25, 60, 100, 36, 66, 83, 16, 5, 26, 96, 37, 99, 93, 10, 52, 88, 19, 2, 87, 55, 47, 7, 28, 91, 21, 35, 85, 62, 31}, {63, 96, 6, 57, 68, 71, 60, 81, 64, 4, 70, 72, 1, 74, 34, 80, 17, 67, 28, 54, 79, 32, 23, 12, 73, 39, 89, 93, 55, 100, 40, 24, 58, 45}, {93, 41, 50, 10, 73, 87, 21, 83, 77, 4, 85, 1, 3, 24, 29, 30, 55, 2, 48, 31, 11, 38, 34, 44, 96, 62, 49, 68, 15, 46, 12, 56, 59, 81}, {48, 47, 87, 10, 30, 78, 23, 44, 11, 58, 39, 7, 9, 53, 20, 46, 73, 77, 40, 69, 2, 62, 21, 88, 75, 43, 90, 14, 36, 13, 85, 91, 54, 28}, {64, 92, 98, 7, 94, 70, 81, 80, 89, 38, 73, 85, 39, 95, 57, 65, 69, 60, 32, 97, 75, 77, 30, 90, 3, 37, 2, 53, 43, 72, 36, 61, 100, 58}, {46, 47, 90, 10, 12, 50, 83, 98, 75, 25, 55, 11, 51, 0, 60, 86, 63, 28, 79, 91, 72, 78, 88, 14, 67, 92, 9, 19, 54, 85, 57, 99, 7, 38}, {13, 14, 28, 69, 34, 1, 97, 23, 53, 91, 9, 75, 58, 61, 21, 24, 50, 95, 79, 93, 52, 26, 72, 86, 89, 40, 96, 65, 90, 45, 15, 87, 18, 63}, {89, 100, 79, 64, 55, 40, 32, 65, 57, 93, 23, 49, 29, 54, 56, 2, 14, 34, 44, 73, 52, 47, 25, 33, 20, 51, 76, 81, 18, 21, 98, 96, 7, 97}, {54, 64, 75, 50, 35, 32, 10, 68, 51, 4, 65, 79, 28, 95, 98, 57, 41, 26, 7, 18, 19, 6, 93, 63, 0, 40, 91, 27, 1, 49, 30, 37, 12, 73}, {83, 73, 24, 70, 8, 72, 6, 21, 80, 93, 43, 13, 82, 33, 99, 37, 26, 40, 69, 64, 11, 2, 81, 53, 46, 38, 55, 76, 62, 74, 92, 49, 30, 4}, {63, 59, 70, 95, 98, 86, 72, 67, 71, 37, 56, 2, 29, 50, 93, 39, 44, 92, 15, 34, 22, 40, 7, 84, 43, 35, 96, 54, 53, 90, 45, 81, 73, 38}, {32, 79, 50, 29, 83, 16, 54, 90, 66, 19, 51, 96, 92, 87, 9, 93, 94, 76, 23, 56, 12, 95, 98, 31, 15, 97, 30, 71, 42, 81, 7, 43, 52, 88}, {65, 53, 23, 63, 82, 81, 9, 67, 75, 44, 39, 69, 16, 55, 99, 93, 17, 71, 36, 96, 84, 85, 28, 19, 35, 79, 70, 1, 25, 26, 72, 80, 22, 10}, {3, 75, 58, 74, 18, 46, 16, 88, 21, 10, 1, 29, 76, 93, 73, 98, 40, 32, 52, 62, 22, 26, 14, 66, 38, 57, 55, 34, 82, 72, 60, 83, 90, 19}, {74, 48, 36, 32, 90, 80, 98, 17, 46, 75, 10, 30, 2, 19, 9, 44, 33, 5, 14, 61, 79, 89, 96, 23, 71, 81, 100, 7, 56, 72, 11, 18, 63, 84}, {93, 81, 35, 27, 18, 95, 13, 44, 23, 49, 54, 86, 87, 96, 19, 8, 57, 24, 62, 66, 74, 3, 32, 78, 46, 17, 22, 53, 30, 56, 25, 6, 61, 9}, {6, 41, 29, 24, 2, 19, 77, 71, 57, 63, 81, 36, 73, 12, 49, 22, 4, 95, 61, 47, 21, 80, 53, 14, 9, 42, 51, 40, 91, 100, 65, 48, 97, 27}, {75, 4, 70, 8, 58, 25, 50, 2, 30, 93, 22, 24, 29, 31, 87, 57, 39, 71, 73, 69, 94, 81, 48, 100, 9, 14, 60, 27, 54, 36, 77, 47, 92, 68}, {15, 48, 8, 83, 28, 41, 61, 52, 3, 23, 0, 67, 55, 87, 60, 73, 46, 35, 47, 25, 71, 95, 69, 31, 54, 1, 76, 10, 94, 13, 58, 14, 51, 93}, {27, 54, 43, 13, 68, 76, 82, 88, 19, 47, 8, 100, 36, 58, 72, 39, 84, 63, 16, 31, 33, 57, 60, 30, 99, 69, 98, 10, 24, 38, 44, 86, 29, 67}, {19, 75, 51, 31, 16, 32, 90, 27, 0, 15, 69, 80, 21, 59, 76, 46, 60, 30, 79, 2, 33, 81, 44, 37, 77, 89, 45, 1, 61, 28, 41, 50, 3, 47}, {72, 19, 99, 80, 33, 86, 2, 96, 100, 22, 85, 74, 55, 18, 42, 70, 21, 73, 76, 95, 90, 26, 51, 75, 38, 41, 3, 69, 49, 31, 7, 61, 68, 53}, {46, 55, 94, 14, 25, 16, 22, 87, 73, 97, 43, 62, 57, 51, 37, 28, 86, 95, 3, 1, 31, 59, 40, 7, 4, 24, 34, 10, 56, 61, 96, 72, 69, 99}, {26, 3, 36, 72, 62, 83, 52, 60, 99, 27, 57, 11, 86, 64, 10, 70, 43, 73, 29, 48, 67, 68, 87, 20, 92, 14, 25, 47, 8, 17, 22, 21, 49, 78}, {12, 3, 10, 76, 97, 87, 57, 30, 42, 20, 9, 73, 27, 23, 51, 35, 92, 19, 46, 68, 70, 24, 61, 85, 11, 77, 72, 59, 5, 95, 80, 67, 50, 52}, {25, 28, 17, 10, 76, 56, 99, 96, 54, 83, 15, 12, 71, 48, 84, 49, 58, 14, 43, 21, 88, 75, 68, 98, 29, 85, 73, 22, 89, 50, 55, 72, 30, 5}, {72, 83, 17, 79, 21, 0, 68, 73, 91, 8, 78, 75, 52, 13, 11, 60, 53, 87, 26, 80, 85, 38, 16, 89, 22, 84, 34, 41, 47, 39, 37, 98, 54, 19}, {42, 39, 16, 8, 44, 36, 96, 63, 79, 83, 5, 33, 98, 7, 64, 35, 49, 10, 41, 40, 18, 73, 65, 70, 23, 47, 55, 50, 28, 4, 26, 3, 100, 13}, {73, 27, 87, 76, 64, 13, 39, 32, 71, 23, 42, 60, 59, 94, 99, 47, 44, 81, 5, 82, 97, 89, 80, 86, 41, 10, 91, 4, 54, 67, 63, 98, 22, 7}, {39, 47, 96, 35, 29, 72, 18, 19, 95, 76, 85, 71, 63, 88, 36, 16, 10, 55, 45, 62, 61, 69, 99, 65, 50, 81, 41, 7, 14, 59, 33, 97, 23, 98}, {56, 38, 52, 81, 50, 42, 57, 26, 82, 72, 58, 10, 77, 25, 73, 13, 69, 6, 7, 30, 74, 70, 0, 34, 41, 29, 84, 17, 1, 55, 66, 9, 60, 44}, {31, 96, 45, 52, 79, 33, 95, 55, 32, 98, 89, 26, 69, 24, 87, 58, 29, 6, 30, 40, 23, 70, 62, 15, 9, 17, 0, 21, 91, 48, 88, 12, 3, 11}, {62, 48, 8, 64, 44, 81, 84, 6, 47, 96, 91, 60, 24, 57, 65, 80, 11, 68, 54, 46, 71, 19, 82, 74, 26, 99, 53, 100, 94, 98, 38, 56, 10, 85}, {77, 97, 12, 72, 92, 73, 6, 85, 9, 95, 11, 83, 54, 51, 61, 66, 20, 16, 39, 90, 93, 5, 8, 47, 75, 32, 22, 35, 33, 60, 56, 18, 38, 0}, {75, 25, 21, 9, 53, 55, 26, 59, 62, 84, 88, 46, 3, 34, 41, 65, 64, 10, 6, 61, 93, 40, 4, 91, 27, 24, 45, 35, 2, 0, 72, 97, 12, 31}, {98, 24, 92, 16, 29, 17, 42, 75, 76, 7, 88, 66, 80, 3, 84, 72, 5, 68, 67, 45, 18, 41, 32, 46, 78, 23, 40, 60, 21, 63, 9, 53, 77, 10}},
				threshold: 51510,
			},
			want: 32,
		},
		{
			name: "maximum side length of a square with sum less than or equal to threshold",
			args: args{
				mat:       [][]int{{93, 76, 38, 49, 22, 87, 0, 65, 94, 96, 18, 82, 97, 58, 39, 29, 90, 33, 23, 95, 61, 88, 89, 16, 73, 15, 84, 51, 37, 57, 64, 54, 50, 77, 42, 6, 5, 35, 8, 1, 100, 45, 67, 24, 31}, {83, 98, 1, 76, 79, 14, 81, 35, 47, 85, 58, 55, 20, 75, 44, 12, 27, 15, 38, 52, 67, 87, 78, 88, 90, 4, 26, 23, 5, 18, 49, 100, 25, 31, 0, 32, 74, 37, 65, 10, 53, 3, 94, 51, 61}, {21, 2, 70, 35, 34, 9, 24, 33, 69, 96, 78, 51, 19, 3, 58, 23, 65, 97, 90, 55, 30, 81, 87, 27, 77, 49, 45, 36, 95, 63, 40, 39, 75, 10, 83, 29, 98, 57, 79, 25, 41, 85, 89, 4, 67}, {60, 71, 61, 36, 32, 58, 47, 37, 84, 77, 48, 16, 42, 75, 20, 70, 68, 69, 2, 22, 72, 98, 15, 80, 62, 9, 63, 59, 12, 97, 30, 78, 3, 91, 5, 39, 86, 56, 34, 66, 81, 29, 67, 93, 55}, {37, 77, 9, 18, 74, 54, 55, 48, 68, 1, 33, 73, 44, 30, 51, 88, 90, 85, 95, 35, 99, 50, 80, 66, 56, 42, 15, 83, 87, 82, 67, 20, 25, 93, 94, 60, 14, 28, 81, 65, 31, 16, 97, 21, 100}, {21, 73, 59, 36, 94, 26, 41, 60, 53, 69, 83, 27, 40, 56, 22, 13, 0, 29, 42, 86, 14, 5, 30, 24, 92, 43, 95, 28, 47, 6, 78, 35, 76, 4, 75, 39, 45, 58, 67, 9, 10, 44, 77, 84, 15}, {10, 27, 50, 33, 20, 94, 60, 95, 19, 78, 35, 43, 65, 88, 2, 28, 92, 45, 91, 9, 3, 4, 84, 86, 34, 81, 82, 89, 46, 100, 73, 98, 36, 62, 42, 18, 16, 69, 21, 90, 11, 12, 1, 70, 49}, {30, 75, 15, 80, 35, 94, 57, 81, 63, 38, 14, 61, 65, 32, 25, 33, 37, 46, 28, 97, 4, 8, 58, 68, 9, 31, 0, 39, 11, 26, 79, 84, 100, 47, 40, 21, 41, 17, 56, 78, 18, 52, 96, 88, 34}, {48, 91, 15, 14, 62, 37, 69, 64, 83, 70, 96, 60, 16, 61, 63, 77, 50, 34, 80, 71, 95, 8, 20, 87, 33, 5, 89, 66, 21, 54, 17, 11, 36, 74, 38, 35, 68, 97, 56, 2, 73, 65, 85, 6, 52}, {38, 3, 65, 48, 29, 95, 37, 32, 69, 40, 34, 25, 71, 63, 44, 13, 30, 7, 11, 4, 52, 88, 45, 21, 6, 33, 28, 72, 23, 31, 27, 97, 5, 93, 58, 77, 18, 74, 85, 35, 56, 20, 78, 62, 79}, {41, 34, 72, 51, 100, 94, 20, 93, 21, 16, 57, 2, 37, 56, 88, 4, 38, 5, 70, 22, 10, 23, 95, 35, 71, 40, 27, 58, 48, 99, 28, 79, 61, 44, 19, 36, 54, 97, 83, 39, 3, 62, 15, 24, 7}, {42, 15, 74, 88, 87, 16, 13, 50, 26, 55, 94, 38, 12, 63, 99, 86, 76, 46, 0, 39, 9, 8, 65, 11, 52, 28, 14, 43, 67, 100, 30, 33, 40, 57, 7, 78, 54, 6, 77, 93, 59, 56, 21, 98, 27}, {5, 100, 70, 0, 14, 13, 3, 30, 53, 31, 69, 91, 79, 66, 64, 59, 26, 82, 77, 27, 15, 61, 57, 51, 47, 6, 45, 83, 73, 63, 34, 39, 44, 87, 43, 8, 85, 99, 80, 48, 71, 1, 96, 55, 11}, {4, 52, 86, 3, 84, 35, 38, 43, 92, 37, 67, 17, 56, 61, 98, 16, 11, 57, 85, 25, 21, 96, 53, 65, 71, 55, 47, 40, 70, 58, 87, 83, 26, 68, 19, 0, 91, 76, 46, 62, 30, 8, 82, 22, 89}, {8, 78, 87, 68, 71, 74, 29, 17, 16, 28, 7, 95, 26, 48, 56, 92, 90, 91, 52, 24, 57, 63, 14, 19, 84, 30, 6, 50, 49, 34, 93, 12, 13, 60, 4, 99, 97, 2, 80, 44, 20, 94, 1, 9, 89}, {69, 0, 65, 85, 19, 9, 72, 14, 43, 18, 28, 73, 78, 87, 67, 27, 59, 63, 53, 2, 89, 96, 8, 26, 84, 91, 22, 42, 4, 61, 71, 54, 1, 17, 49, 47, 100, 24, 46, 93, 70, 48, 64, 99, 56}, {98, 36, 78, 44, 8, 99, 71, 86, 87, 69, 60, 61, 48, 39, 35, 32, 7, 100, 29, 19, 15, 85, 66, 9, 41, 59, 22, 13, 79, 11, 77, 12, 0, 25, 93, 42, 33, 53, 2, 6, 94, 64, 95, 73, 16}, {62, 12, 30, 100, 90, 81, 23, 72, 39, 10, 13, 78, 74, 15, 18, 4, 76, 47, 73, 70, 24, 19, 45, 61, 63, 65, 99, 60, 98, 20, 5, 34, 52, 14, 95, 54, 66, 55, 91, 94, 71, 83, 2, 50, 59}, {7, 38, 29, 97, 98, 66, 3, 39, 55, 74, 64, 41, 30, 60, 6, 73, 100, 53, 52, 84, 67, 78, 47, 72, 63, 46, 43, 40, 5, 56, 62, 59, 36, 94, 51, 54, 37, 80, 33, 17, 34, 75, 20, 50, 68}, {69, 35, 91, 67, 82, 0, 73, 94, 98, 40, 56, 48, 22, 17, 4, 14, 24, 64, 34, 84, 92, 55, 5, 88, 45, 31, 93, 96, 54, 59, 26, 60, 11, 53, 50, 58, 65, 12, 25, 18, 77, 74, 10, 72, 20}, {69, 49, 81, 100, 5, 58, 21, 82, 23, 85, 7, 4, 52, 11, 26, 86, 1, 63, 41, 34, 13, 65, 68, 43, 12, 78, 32, 79, 93, 0, 83, 15, 27, 54, 47, 42, 97, 39, 72, 33, 60, 64, 88, 45, 3}, {7, 29, 10, 63, 33, 51, 47, 90, 98, 99, 28, 75, 41, 12, 56, 91, 88, 62, 68, 52, 13, 16, 20, 77, 19, 94, 17, 6, 43, 67, 26, 69, 37, 89, 34, 40, 49, 64, 78, 46, 42, 27, 93, 57, 38}, {80, 75, 2, 40, 67, 49, 99, 73, 52, 68, 33, 53, 20, 96, 54, 34, 10, 58, 38, 86, 26, 92, 24, 56, 74, 47, 14, 89, 4, 81, 23, 5, 84, 46, 27, 91, 70, 55, 35, 16, 87, 7, 97, 29, 22}, {3, 2, 58, 36, 80, 43, 8, 42, 46, 84, 17, 6, 100, 51, 66, 95, 69, 18, 15, 53, 57, 47, 52, 12, 76, 71, 70, 40, 25, 37, 87, 29, 41, 81, 39, 26, 89, 48, 33, 24, 55, 22, 79, 63, 31}, {63, 41, 45, 20, 35, 77, 78, 57, 33, 62, 48, 80, 84, 91, 10, 22, 19, 40, 53, 55, 65, 59, 83, 23, 94, 24, 27, 60, 11, 56, 92, 73, 97, 64, 17, 47, 13, 44, 4, 38, 39, 75, 58, 76, 89}, {49, 69, 55, 90, 72, 28, 81, 32, 47, 91, 98, 29, 30, 67, 20, 26, 78, 89, 31, 77, 12, 83, 99, 23, 43, 1, 11, 68, 56, 38, 61, 62, 15, 74, 76, 73, 79, 9, 0, 25, 65, 19, 57, 96, 50}, {67, 20, 16, 4, 42, 55, 13, 10, 36, 41, 77, 29, 91, 87, 40, 21, 93, 44, 56, 92, 1, 9, 12, 46, 68, 7, 81, 6, 84, 99, 66, 27, 95, 19, 100, 35, 2, 72, 98, 43, 75, 0, 47, 74, 31}, {63, 64, 4, 22, 96, 50, 0, 23, 60, 8, 80, 1, 56, 5, 82, 36, 73, 92, 67, 35, 65, 84, 79, 83, 46, 37, 9, 14, 54, 15, 71, 24, 13, 53, 99, 3, 58, 55, 21, 29, 25, 2, 87, 59, 51}, {51, 90, 24, 37, 49, 68, 20, 83, 13, 21, 65, 73, 15, 47, 17, 30, 57, 38, 26, 40, 8, 82, 92, 35, 10, 29, 44, 81, 56, 100, 59, 6, 88, 89, 16, 39, 1, 93, 76, 71, 4, 61, 97, 33, 32}, {90, 2, 7, 73, 98, 67, 100, 70, 93, 22, 71, 66, 92, 96, 57, 41, 99, 74, 21, 32, 34, 10, 5, 42, 95, 58, 1, 30, 0, 49, 69, 54, 79, 46, 44, 18, 97, 13, 3, 23, 89, 35, 68, 4, 6}, {20, 25, 78, 28, 90, 11, 18, 24, 66, 65, 58, 1, 100, 3, 51, 79, 91, 12, 83, 98, 67, 96, 39, 86, 82, 41, 16, 40, 43, 32, 38, 23, 63, 97, 5, 73, 56, 70, 19, 71, 75, 92, 15, 60, 4}, {29, 13, 26, 48, 22, 40, 49, 30, 89, 27, 95, 19, 35, 83, 8, 47, 9, 78, 80, 82, 7, 21, 32, 2, 72, 71, 62, 38, 90, 6, 45, 53, 31, 59, 52, 50, 92, 34, 4, 42, 99, 87, 65, 33, 41}, {23, 17, 29, 55, 84, 70, 73, 10, 37, 88, 69, 11, 71, 45, 1, 52, 85, 5, 3, 67, 64, 25, 62, 74, 20, 54, 42, 15, 65, 80, 92, 57, 60, 86, 98, 16, 100, 36, 93, 89, 50, 63, 79, 61, 47}, {73, 99, 4, 24, 29, 100, 66, 92, 67, 32, 0, 1, 25, 23, 59, 62, 56, 44, 22, 58, 34, 9, 46, 33, 75, 71, 5, 6, 64, 95, 41, 48, 19, 93, 45, 15, 39, 94, 2, 55, 49, 53, 78, 31, 72}, {2, 34, 83, 43, 75, 73, 51, 13, 60, 57, 5, 81, 96, 90, 28, 22, 59, 19, 61, 24, 20, 49, 27, 53, 9, 39, 45, 8, 7, 70, 36, 35, 66, 15, 6, 58, 91, 52, 98, 18, 14, 17, 67, 87, 74}, {33, 61, 10, 6, 52, 70, 53, 11, 39, 75, 4, 22, 47, 36, 79, 74, 67, 97, 65, 72, 57, 86, 91, 21, 17, 83, 16, 3, 19, 90, 48, 60, 0, 45, 98, 63, 20, 87, 66, 64, 32, 1, 56, 73, 30}, {25, 73, 43, 20, 17, 3, 100, 15, 49, 69, 37, 99, 11, 48, 21, 24, 44, 78, 5, 83, 98, 29, 2, 34, 14, 10, 38, 56, 62, 40, 36, 6, 1, 23, 87, 60, 4, 31, 86, 94, 66, 30, 74, 77, 59}, {85, 18, 37, 10, 2, 79, 57, 60, 11, 48, 22, 81, 44, 19, 62, 95, 5, 7, 69, 15, 46, 52, 3, 41, 39, 75, 88, 31, 67, 100, 89, 82, 74, 87, 92, 93, 58, 66, 6, 86, 72, 77, 80, 99, 64}, {30, 18, 40, 70, 90, 83, 75, 52, 45, 89, 5, 85, 13, 73, 7, 6, 94, 15, 74, 92, 2, 59, 37, 84, 38, 54, 55, 93, 66, 8, 51, 87, 10, 86, 1, 28, 49, 41, 64, 100, 67, 68, 65, 26, 95}, {99, 51, 66, 37, 41, 69, 23, 15, 89, 26, 35, 39, 17, 10, 72, 48, 5, 18, 28, 63, 76, 50, 55, 42, 33, 68, 97, 59, 7, 2, 8, 40, 30, 74, 100, 24, 91, 31, 70, 64, 0, 47, 3, 56, 85}, {7, 77, 22, 87, 84, 72, 20, 89, 47, 88, 78, 85, 64, 82, 94, 95, 0, 34, 73, 66, 24, 50, 42, 30, 9, 6, 12, 75, 69, 38, 1, 2, 28, 40, 62, 21, 83, 67, 35, 53, 61, 86, 60, 59, 18}, {8, 41, 62, 99, 23, 92, 89, 22, 38, 7, 56, 86, 100, 63, 73, 77, 74, 98, 65, 6, 37, 43, 84, 35, 50, 95, 61, 94, 42, 10, 16, 4, 83, 60, 51, 96, 75, 34, 67, 0, 48, 54, 66, 33, 69}, {82, 77, 8, 12, 76, 69, 96, 47, 7, 10, 9, 22, 87, 93, 43, 33, 48, 90, 63, 45, 32, 26, 79, 51, 3, 98, 44, 46, 86, 97, 94, 4, 72, 1, 14, 60, 66, 11, 56, 34, 61, 31, 25, 53, 23}, {83, 87, 42, 21, 20, 34, 98, 0, 58, 47, 100, 40, 33, 99, 66, 45, 69, 37, 52, 61, 24, 88, 78, 14, 90, 76, 68, 79, 23, 29, 73, 81, 26, 50, 32, 5, 63, 2, 18, 35, 8, 59, 94, 3, 67}, {62, 66, 88, 49, 35, 13, 24, 29, 90, 51, 98, 95, 14, 57, 20, 74, 21, 69, 91, 53, 60, 87, 56, 86, 33, 54, 10, 100, 12, 75, 39, 45, 9, 3, 94, 85, 23, 30, 47, 46, 55, 99, 50, 7, 43}, {6, 24, 83, 20, 67, 59, 48, 21, 81, 11, 99, 87, 95, 17, 46, 96, 18, 58, 77, 93, 89, 26, 37, 86, 22, 100, 56, 79, 90, 27, 44, 8, 78, 28, 33, 35, 75, 25, 74, 88, 15, 40, 49, 30, 84}},
				threshold: 4749,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxSideLength(tt.args.mat, tt.args.threshold))
		})
	}
}
