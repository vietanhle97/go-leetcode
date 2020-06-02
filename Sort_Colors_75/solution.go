package Sort_Colors_75

func partition(low, high int, nums []int) int {
	i := low - 1
	pivot := nums[high]
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i += 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

func quickSort(low, high int, nums []int) {
	if low < high {
		pivot := partition(low, high, nums)
		quickSort(low, pivot-1, nums)
		quickSort(pivot+1, high, nums)
	}
}

func sortColors(nums []int) {
	quickSort(0, len(nums)-1, nums)
}
