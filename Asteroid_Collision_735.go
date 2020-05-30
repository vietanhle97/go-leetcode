package main

import "fmt"

func last(nums []int) int {
	return nums[len(nums)-1]
}

func isEmpty(nums []int) bool {
	return len(nums) == 0
}
func asteroidCollision(asteroids []int) []int {
	m := len(asteroids)
	if m < 2 {
		return asteroids
	}

	ind, first := -1, -1
	res := []int{}
	for i := 0; i < m; i++ {
		if asteroids[i] > 0 {
			ind, first = i, asteroids[i]
			break
		} else {
			res = append(res, asteroids[i])
		}
	}

	if ind == -1 && first == -1 {
		return asteroids
	}

	pos := []int{}
	for i := ind; i < m; i++ {
		cur := asteroids[i]
		if cur > 0 {
			pos = append(pos, cur)
		} else {
			if isEmpty(pos) {
				res = append(res, cur)
				continue
			}
			equal := false
			for last(pos) <= -cur {
				if last(pos) == -cur {
					equal = true
					pos = pos[:len(pos)-1]
					break
				}
				pos = pos[:len(pos)-1]
				if isEmpty(pos) {
					break
				}
			}

			if !equal && isEmpty(pos) {
				res = append(res, cur)
			}
		}
	}

	return append(res, pos...)
}

func main() {
	a := []int{-8, -8, 7, -8}
	fmt.Println(asteroidCollision(a))

}
