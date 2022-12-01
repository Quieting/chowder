package copier

import (
	"testing"

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
