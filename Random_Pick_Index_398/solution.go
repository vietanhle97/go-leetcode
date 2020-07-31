package Random_Pick_Index_398

import (
	"math/rand"
	"time"
)

type Solution struct {
	n   map[int][]int
	cnt int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().UnixNano())
	n := map[int][]int{}
	cnt := 0
	for i := 0; i < len(nums); i++ {
		n[nums[i]] = append(n[nums[i]], cnt)
		cnt++
	}
	// fmt.Println(n)
	return Solution{n, cnt}
}

func (random *Solution) Pick(target int) int {
	// fmt.Println(random.n[target])
	pick := rand.Intn(len(random.n[target]))
	return random.n[target][pick]
}
