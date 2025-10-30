func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}	
		}
	}
	return []int{}
}



//更优解
func twoSum1(nums []int, target int) []int {
    set := make(map[int]int)
    for i, num := range nums {
        if v, ok := set[target-num]; ok {
            return []int{v, i}
        }
        set[num] = i
    }
    return []int{-1, -1}
}
