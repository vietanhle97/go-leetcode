package Range_Sum_Query_Mutable_307

type NumArray struct {
	A   []int
	bit []int
}

func count(bit []int, i int) int {
	i++
	s := 0
	for i > 0 {
		s += bit[i]
		i -= i & (-i)
	}
	return s
}

func update(bit []int, n, i, v int) {
	i++
	for i <= n {
		bit[i] += v
		i += i & (-i)
	}
}

func Constructor(nums []int) NumArray {
	bit := make([]int, len(nums)+1)
	for i, e := range nums {
		update(bit, len(nums), i, e)
	}
	return NumArray{nums, bit}
}

func (BIT *NumArray) Update(i int, val int) {
	tmp := BIT.A[i]
	BIT.A[i] = val
	update(BIT.bit, len(BIT.A), i, BIT.A[i]-tmp)

}

func (BIT *NumArray) SumRange(i int, j int) int {
	return count(BIT.bit, j) - count(BIT.bit, i) + BIT.A[i]
}
