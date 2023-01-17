func binary_search(nums []int, target int, 
    low_index int, high_index int) int {

    if low_index > high_index {
        return -1
    }
    mid := (high_index + low_index) / 2
    if nums[mid] == target {
        return mid
    } else if nums[mid] < target {
        return binary_search(nums, target, mid + 1, high_index)
    } else {
        return binary_search(nums, target, low_index, mid - 1)
    }

}

func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    
    return binary_search(nums, target, 0, len(nums) - 1)
}
