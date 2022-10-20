package copier

import (
	"testing"

	"github.com/jinzhu/copier"
)

var rsc = A{
	age: 20,
}
var dst = B{
	age: 10,
}

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Copy(&rsc, &dst)
	}
}

func BenchmarkCopier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copier.Copy(&rsc, &dst)
	}
}
