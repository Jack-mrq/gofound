package arrays

const (
	LOW  = 0
	HIGH = 1
)

// BinarySearch 二分查找
func BinarySearch(arr []uint32, target uint32) bool {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

//遍历字符串数组
func ArrayStringExists(arr []string, str string) bool {

	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false

}

// MergeArrayUint32 合并两个数组
//question ? source 为什么加在后面
func MergeArrayUint32(target []uint32, source []uint32) []uint32 {

	for _, val := range source {
		if !BinarySearch(target, val) {
			target = append(target, val)
		}
	}
	return target
}

//找到对应值，并返回下标
func Find(arr []uint32, target uint32) int {
	for index, v := range arr {
		if v == target {
			return index
		}
	}
	return -1
}
