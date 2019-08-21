package snake

// TwoDimensional 初始化的二维切片
func TwoDimensional(width, height uint16) [][]bool {
	s := make([][]bool, 0, height)
	for i := uint16(0); i < height; i++ {
		s = append(s, make([]bool, width, width))
	}
	return s
}
