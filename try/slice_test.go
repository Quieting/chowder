package try

import "testing"

// 测试空切片是否可以使用下标赋值
// 结果：空切片不允许使用下标赋值，会抛出切片越界panic
func Test_Empty_Slice(t *testing.T) {
	var val []int
	val[0] = 10
}
