package main

func removeDuplicates(nums []int) int {//[0,1,1,1,2,2,3,3,4] 
	if len(nums) == 0 {
		return 0
	}
	//var k int = 0
	for i := 0; i < len(nums) -1; i++ {//i=1 i=2 相等
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}


//更优解法
func removeDuplicates1(nums []int) int {//[0,1,1,1,2,2,3,3,4] 
	if len(nums) == 0 {
		return 0
	}
		//不重复的元素个数，第一个元素必然不重复
		slow:=1;
	//var k int = 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast];
			slow++;
		}
	}
	return slow
}
