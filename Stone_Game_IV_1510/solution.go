package Stone_Game_IV_1510

func winnerSquareGame(n int) bool {
	res := make([]bool, n+1)

	for i := 1; i < n+1; i++ {
		for j := 1; j*j < i+1; j++ {
			tmp := j * j
			ls := i - tmp
			if !res[ls] {
				res[i] = true
				break
			}
		}
	}
	return res[n]
}
