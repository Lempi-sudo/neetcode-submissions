func hasDuplicate(nums []int) bool {
    freq:= make(map[int]struct{})
    for _, num := range nums{
        if _,ok:= freq[num]; ok{
            return true
        }else{
            freq[num]=struct{}{}
        }
    }
    return false
}
