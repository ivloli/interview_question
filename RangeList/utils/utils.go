package utils

//LeftBound 获取大于等于目标的最小index
func LeftBound(s []int, start, end, target int) (int, bool) {
	i, j := start, end
	for i < j {
		mid := i + (j-i)>>1
		if s[mid] >= target {
			j = mid
		} else {
			i = mid + 1
		}
	}
	if i < len(s) {
		f := s[i] == target
		return i, f
	}
	return -1, false
}

//RightBound 获取小于等于目标的最大index
func RightBound(s []int, start, end, target int) (int, bool) {
	i, j := start, end
	for i < j {
		mid := i + (j-i)>>1
		if s[mid] <= target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	var f bool
	if j > 0 {
		f = (s[j-1] == target)
	}
	return j - 1, f
}

//IsEven 判断是否是偶数
func IsEven(a int) bool {
	if (a>>1)<<1 == a {
		return true
	}
	return false
}
