package Asteroid_Collision_735

// Time complexity: O(n)
// Space complexity: O(n)

// Get the last element of the array
func last(nums []int) int {
	return nums[len(nums)-1]
}

// Check the array is empty
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
			// Find the first positive asteroid (index and value)
			ind, first = i, asteroids[i]
			break
		} else {
			// Add all the negative asteroids to the result before meeting the first positive asteroid
			res = append(res, asteroids[i])
		}
	}
	// If no positive asteroids found => return the given asteroids list
	if ind == -1 && first == -1 {
		return asteroids
	}
	//Initialize the stack
	pos := make([]int, 0)
	// Go from the position of the first positive asteroid
	for i := ind; i < m; i++ {
		cur := asteroids[i]
		if cur > 0 {
			// whenever meeting a positive asteroid => add to the stack
			pos = append(pos, cur)
		} else {
			// when meet the negative and stack is empty => add the negative to the result
			if isEmpty(pos) {
				res = append(res, cur)
				continue
			}
			sumToZero := false
			// pop the stack until finding the positive asteroid that is bigger than or equal
			// to the absolute value of the negative
			for last(pos) <= -cur {
				// if equal => pop it and break
				if last(pos) == -cur {
					sumToZero = true
					pos = pos[:len(pos)-1]
					break
				}
				// else => just pop the stack
				pos = pos[:len(pos)-1]
				// if stack empty => break
				if isEmpty(pos) {
					break
				}
			}
			// only add the negative if stack empty and we found no element that sum to zero
			if !sumToZero && isEmpty(pos) {
				res = append(res, cur)
			}
		}
	}
	// append the res of stack to the result
	return append(res, pos...)
}
