package Smallest_Integer_Divisible_By_K_1015

type Node struct {
	r int
	s string
}

func smallestRepunitDivByKSol2(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}
	if K == 1 {
		return 1
	}
	vis := make([]bool, K)
	q := []*Node{&Node{1, "1"}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		vis[cur.r] = true
		if vis[(cur.r*10+1)%K] {
			break
		}
		if (cur.r*10+1)%K == 0 {
			return len(cur.s) + 1
		} else {
			q = append(q, &Node{(cur.r*10 + 1) % K, cur.s + "1"})
		}
	}
	return -1
}
