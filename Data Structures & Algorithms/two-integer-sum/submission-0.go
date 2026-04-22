func twoSum(nums []int, target int) []int {
		d:= make(map[int]int)
		for i,n:=range nums{
			x:= target- n
			if ind,ok:= d[x];ok{
				return []int{ind,i}
			}else{
				d[n]=i
			}
		}
	
	return []int{}
}
