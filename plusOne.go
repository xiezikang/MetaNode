func plusOne(digits []int) []int { //9
	n := len(digits) //1
	for i := n-1; i >=0; i-- {
		if digits[i] + 1 ==10 {
			digits[i] =0;
			if(i==0){
				digits =append([]int{1},digits...)
			}
			continue
		}
		digits[i]++
		return digits;

		
	}
	return digits;
}
