package md

import (
	"reflect"
	"testing"
	"unsafe"
)

// growSlice 切片扩容新容量计算规则
func growSlice(oldCap, cap int) int {
	newCap := oldCap
	doubleCap := newCap + newCap
	if cap > doubleCap {
		newCap = cap
	} else {
		const threshold = 256
		if oldCap < threshold {
			newCap = doubleCap
		} else {
			for 0 < newCap && newCap < cap {
				newCap += (newCap + 3*threshold) / 4
			}
			if newCap < 0 {
				newCap = cap
			}
		}
	}

	return newCap
}

func TestGrowSlice(t *testing.T) {
	args := []struct {
		name        string
		oldCap, cap int
	}{
		{name: "增长容量较小", oldCap: 10, cap: 12},
		{name: "原大小较小，增长较大", oldCap: 10, cap: 10000},
		{name: "原容量较大，增长较小", oldCap: 10000, cap: 13200},
	}

	for _, arg := range args {
		t.Run(arg.name, func(t *testing.T) {
			t.Logf("%d\n", growSlice(arg.oldCap, arg.cap))
		})
	}
}

// TestStringToByte 测试字符串转字节切片有没有将结束标记'0'带出来
func TestStringToByte(t *testing.T) {
	s := "HELLO"
	t.Logf("s(字符串)长度：%d\n", len(s))
	t.Logf("s(字符串)内容：%s\n", s)

	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	var b []byte
	pBytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pBytes.Data = stringHeader.Data
	pBytes.Len = stringHeader.Len
	pBytes.Cap = stringHeader.Len
	t.Logf("b(切片)长度：%d\n", len(b))
	t.Logf("b(切片)内容：%s\n", b)

	var b1 = make([]byte, len(b))
	_ = copy(b1, b)
	t.Logf("b1(切片)长度：%d\n", len(b1))
	t.Logf("b1(切片)内容：%s\n", string(b1))

	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b1))

	var s1 string
	pStr := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	pStr.Data = bytesHeader.Data
	pStr.Len = bytesHeader.Len

	t.Logf("s1(字符串)长度：%d\n", len(s1))
	t.Logf("s1(字符串)内容：%s\n", s1)
}
