package Number_Of_Burgers_With_No_Waste_Of_Ingredients_1276

func numOfBurgers(T, C int) []int {
	if T%2 == 1 {
		return []int{}
	}
	a, b := (T-2*C)/2, (4*C-T)/2
	if a >= 0 && b >= 0 {
		return []int{a, b}
	}
	return []int{}
}
