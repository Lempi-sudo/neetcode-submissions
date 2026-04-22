func twoSum(numbers []int, target int) []int {
		l ,r := 0 , len(numbers)-1

        for l <=r{
			cur := numbers[l]  + numbers[r]
			if cur == target{
				return []int{l+1,r+1}
			}else if cur>target{
				r-=1
			}else{
				l+=1
			}
		}
		return []int{}
}
