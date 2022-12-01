package md

import (
	"testing"
	"unsafe"
)

func TestChan(t *testing.T) {
	// var ch = make(chan int, 2)
	//// 先读后写
	//t.Log(<-ch)
	//go func(ch chan<- int) {
	//	ch <- 1
	//	close(ch)
	//}(ch)
	//
	//// ch := &ch{
	//// 	ch: make(chan int, 2),
	//// }
	//// ch.reserve()
	//// go ch.send()
	//
	//time.Sleep(1 * time.Second)

	var i = make(map[string]struct {
		Name string
	})
	t.Logf("%v", &i)
}

func TestSlice(t *testing.T) {
	//a := []string{"a", "a", "a", "a"}
	a := make([]string, 0, 5)
	a = append(a, "a", "a", "a", "a")
	a1 := append(a[:2], []string{"b", "b"}...)
	a2 := append(a[1:2], []string{"c", "c"}...)
	a3 := append(a[2:], []string{"d", "d"}...)
	a4 := a[2:]
	_ = a4
	t.Logf("a:%+v, a1:%+v,a2:%v,a3:%v", a, a1, a2, a3)
}

// TestStringToBytes string 和 []byte 互转相关测试
// 1。string 底层记录的是字节长度而不是字符长度
// 2  []rune 和 []int32 完全一样
func TestStringToBytes(t *testing.T) {
	s := "golang 大法好！"
	bs := *(*[]byte)(unsafe.Pointer(&s))

	t.Logf("字符串长度：%d，字节切片长度：%d\n", len(s), len(bs))
	t.Logf("字符长度：%d，字节切片长度：%d\n", len([]rune(s)), len([]int32(s)))

	for i, val := range bs {
		t.Logf("%d: %b\n", i, val)
	}

}

// 测试创建 chan 时的各种细节
func TestMakeChan(t *testing.T) {
	// chan 类型包含指针变量
	type A struct {
		Name string
		Age  int64
	}
	type Ptr struct {
		// FieldA  *A
		// FieldA1 *A
		// Field1  string
		Field2 int
	}

	var ch = make(chan Ptr, 2)

	t.Log(7 &^ 4)
	t.Log(4 << (^uintptr(0) >> 63))

	ch <- Ptr{}
}

type ch struct {
	ch chan int
}

func (ch *ch) send() {
	ch.ch <- 6
}

func (ch *ch) reserve() {
	_ = <-ch.ch
}

func TestMap(t *testing.T) {
	var m = make(map[string]int32, 1024)
	m["age"] = 18
	m["height"] = 175
	m["weight"] = 70

	for key, val := range m {
		t.Logf("key: %s; val: %d\n", key, val)
	}
	var B uint8
	var hint = 1 << 32
	for overLoadFactor(hint, B) {
		B++
	}
	const deBruijn64ctz = 0x0218a392cd3d5dbf

	t.Logf("B:%d\n", B)
	t.Logf("_PageSize(1<<13):%d\n", 1<<13)
	t.Logf("divRoundUp:%d\n", divRoundUp(12, 8))
	t.Logf("alignUp:%d\n", alignUp(12, 8))
	n := uint64(-16 & 16)
	t.Logf("正负数与运算结果: %d\n", n)
	t.Logf("正负数与运算结果: %x\n", (n)*deBruijn64ctz)
	t.Logf("正负数与运算结果: %d\n", (n)*deBruijn64ctz)
	t.Logf("正负数与运算结果: %x\n", n*deBruijn64ctz>>58)

	_ = new(int64)
}

func bucketShift(b uint8) uintptr {
	// Masking the shift amount allows overflow checks to be elided.
	return uintptr(1) << (b & (8*8 - 1))
}
func overLoadFactor(count int, B uint8) bool {
	return count > 8 && uintptr(count) > 13*(bucketShift(B)/2)
}

// divRoundUp returns ceil(n / a).
func divRoundUp(n, a uintptr) uintptr {
	// a is generally a power of two. This will get inlined and
	// the compiler will optimize the division.
	return (n + a - 1) / a
}

// alignUp rounds n up to a multiple of a. a must be a power of 2.
func alignUp(n, a uintptr) uintptr {
	return (n + a - 1) &^ (a - 1)
}
