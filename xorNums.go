func xorAllNums(nums1 []int, nums2 []int) int {
    var x, y int
    x = 0
    y = 0

    for i := range nums1 {
        x ^= nums1[i]
    }
    for j := range nums2 {
        y ^= nums2[j]
    }
    return (len(nums1) % 2 * y) ^ (len(nums2) % 2 * x)
}
