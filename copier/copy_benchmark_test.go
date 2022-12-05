package copier

import (
	"testing"
	"unsafe"

	"github.com/jinzhu/copier"
)

func BenchmarkCopy(b *testing.B) {
	rsc := A{
		age: 20,
	}
	dst := B{
		age: 10,
	}
	for i := 0; i < b.N; i++ {
		Copy(&rsc, &dst)
	}
}

func BenchmarkMemove(b *testing.B) {
	from, to := A{age: 20}, A{}
	for i := 0; i < b.N; i++ {
		memove(uintptr(unsafe.Pointer(&from.age)), uintptr(unsafe.Pointer(&to.age)), 8)
	}
}

func BenchmarkMemove1(b *testing.B) {
	from, to := A{age: 20}, A{}
	for i := 0; i < b.N; i++ {
		memove1(uintptr(unsafe.Pointer(&from.age)), uintptr(unsafe.Pointer(&to.age)), 8)
	}
}

func BenchmarkCopier(b *testing.B) {
	rsc := A{
		age: 20,
	}
	dst := B{
		age: 10,
	}
	for i := 0; i < b.N; i++ {
		copier.Copy(&rsc, &dst)
	}
}
