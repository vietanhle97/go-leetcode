package Random_Pick_With_Weight_528

import (
	"math/rand"
	"time"
)

type Solution struct {
	sums []int
}

func Constructor(w []int) Solution {
	rand.Seed(time.Now().UnixNano())
	sums := make([]int, len(w))
	sums[0] = w[0]
	for i := 1; i < len(w); i++ {
		sums[i] = w[i] + sums[i-1]
	}

	return Solution{sums}
}

func (random *Solution) PickIndex() int {
	pick := rand.Intn(random.sums[len(random.sums)-1])
	l, r := 0, len(random.sums)-1
	for l <= r {
		mid := l + (r-l)/2
		if random.sums[mid] > pick {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
