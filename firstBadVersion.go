func binarySearchHelper(low int, high int) int {
    if low == high {
        return low
    }
    mid := (high + low) / 2
    if isBadVersion(mid) {
        if mid > 1 && isBadVersion(mid - 1) {
            return binarySearchHelper(low, mid - 1)
        } else if mid > 1 && !isBadVersion(mid - 1) {
            return mid
        }
    } else {
        return binarySearchHelper(mid + 1, high)
    }
    return low
}

func firstBadVersion(n int) int {
    return binarySearchHelper(1, n)
}
