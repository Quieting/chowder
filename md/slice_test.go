package md

import (
	"testing"
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
