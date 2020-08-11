package Design_A_Stack_With_Increment_Operation_1381

type CustomStack struct {
	max int
	A   []int
}

func Constructor(maxSize int) CustomStack {
	A := make([]int, 0)
	return CustomStack{maxSize, A}
}

func (st *CustomStack) Push(x int) {
	if len(st.A) == st.max {
		return
	}
	st.A = append(st.A, x)
}

func (st *CustomStack) Pop() int {
	if len(st.A) == 0 {
		return -1
	}
	ans := st.A[len(st.A)-1]
	st.A = st.A[:len(st.A)-1]
	return ans
}

func (st *CustomStack) Increment(k int, val int) {
	i := 0
	for i < k {
		if i >= len(st.A) {
			break
		}
		st.A[i] = st.A[i] + val
		i++
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
