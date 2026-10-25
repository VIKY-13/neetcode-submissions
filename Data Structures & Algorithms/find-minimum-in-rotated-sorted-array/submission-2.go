func findMin(nums []int) int {
    low, high := 0, len(nums)-1
    res := nums[0]

    for low <= high{
        mid := low +(high-low)/2

        if nums[low] <= nums[mid]{
            if nums[low] < res{
                res = nums[low]
            }
            low = mid+1
        }else{
            if nums[mid] < res{
                res = nums[mid]
            }
            high = mid-1
        }
    }

    return res
}
