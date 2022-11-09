package rangeList

import (
	"core/utils"
	"fmt"
	"strconv"
	"strings"
)

// RangeList 底层为一个数组, 根据index 偶数为左边界,奇数为右边界
// example: []int{1,2,3,4,5,6} -> [1,2) [3,4) [5,6)
type RangeList struct {
	interSlice []int
}

//LeftBound 获取左边界
func (s *RangeList) LeftBound(start, end, target int) (int, bool) {
	return utils.LeftBound(s.interSlice, start, end, target)
}

//RightBound 获取右边界
func (s *RangeList) RightBound(start, end, target int) (int, bool) {
	return utils.RightBound(s.interSlice, start, end, target)
}

//checkInput 检查入参
func (s *RangeList) checkInput(left, right int) bool {
	if left >= right {
		return false
	}
	return true
}

//RemoveRange 移除一个range
func (s *RangeList) RemoveRange(left, right int) {
	if !s.checkInput(left, right) {
		return
	}
	if s.interSlice == nil {
		s.interSlice = []int{}
		return
	}

	//寻找range左值在rangeList中的左边界 没有表示目标range大于所有现有range
	lb, _ := s.LeftBound(0, len(s.interSlice), left)
	if lb == -1 {
		return
	}
	//寻找range右值在rangeList中右边界 没有表示目标range小于所有现有range
	rb, _ := s.RightBound(lb, len(s.interSlice), right)
	if rb == -1 {
		return
	}

	// left [..) right or [..) left right [..)
	if utils.IsEven(lb) && !utils.IsEven(rb) {
		s.interSlice = append(s.interSlice[0:lb], s.interSlice[rb+1:]...)
		return
	}
	// [..left right ..)
	if lb > rb {
		s.interSlice = append(s.interSlice[0:lb], append([]int{left, right}, s.interSlice[rb+1:]...)...)
		return
	}
	mergeLeft, mergeRight := lb, rb
	if !utils.IsEven(lb) {
		mergeLeft++
		s.interSlice[lb] = left
	}
	if utils.IsEven(rb) {
		mergeRight--
		s.interSlice[rb] = right
	}
	s.interSlice = append(s.interSlice[0:mergeLeft], s.interSlice[mergeRight+1:]...)
}

//AddRange 添加一个range
func (s *RangeList) AddRange(left, right int) {
	if !s.checkInput(left, right) {
		return
	}
	if s.interSlice == nil {
		s.interSlice = []int{left, right}
		return
	}
	lb, _ := s.LeftBound(0, len(s.interSlice), left)
	if lb == -1 {
		s.interSlice = append(s.interSlice, left, right)
		return
	}
	rb, _ := s.RightBound(lb, len(s.interSlice), right)
	if rb == -1 {
		s.interSlice = append([]int{left, right}, s.interSlice...)
		return
	}
	//  left [..) right or [..) left right [..)
	if utils.IsEven(lb) && !utils.IsEven(rb) {
		s.interSlice = append(s.interSlice[0:lb], append([]int{left, right}, s.interSlice[rb+1:]...)...)
		return
	}
	mergeLeft, mergeRight := lb, rb
	if utils.IsEven(lb) {
		s.interSlice[lb] = left
		mergeLeft++
	}
	if !utils.IsEven(rb) {
		s.interSlice[rb] = right
		mergeRight--
	}
	//剪枝优化 [..left..) right or left [..right..)
	if lb == rb {
		return
	}
	s.interSlice = append(s.interSlice[0:mergeLeft], s.interSlice[mergeRight+1:]...)
}

//QueryRange 查询一个range是否被完全包含
func (s *RangeList) QueryRange(left, right int) bool {
	if left >= right {
		return true
	}
	lb, eql := s.LeftBound(0, len(s.interSlice), left)
	if lb == -1 {
		return false
	}
	rb, eqr := s.RightBound(lb, len(s.interSlice), right)
	if rb == -1 {
		return false
	}
	if eql {
		lb++
	}
	if eqr {
		rb--
	}
	if lb > rb && utils.IsEven(rb) && !utils.IsEven(lb) {
		return true
	}
	return false

}

//ToString 将rangeList转换为可读string: [1,2) [3,4)...
func (s *RangeList) ToString() string {
	if len(s.interSlice) == 0 {
		return "[)"
	}
	if len(s.interSlice)%2 != 0 {
		return "Data Error"
	}
	StrBuilder := strings.Builder{}
	for i := 0; i < len(s.interSlice); i += 2 {
		StrBuilder.WriteString("[" + strconv.Itoa(s.interSlice[i]) + "," + strconv.Itoa(s.interSlice[i+1]) + ")")
		if i < len(s.interSlice)-2 {
			StrBuilder.WriteByte(' ')
		}
	}
	return StrBuilder.String()
}

//Print 打印rangeList的可读字符串
func (s *RangeList) Print() {
	fmt.Println(s.ToString())
}
