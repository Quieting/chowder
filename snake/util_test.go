package snake

import (
	"testing"
)

func TestTwoDimensional(t *testing.T) {
	s := TwoDimensional(3, 3)
	t.Logf("s[2][2] = %t\n", s[2][2])
}
