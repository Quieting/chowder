package math

// MatrixMulti 矩阵乘法
// 需要保证 矩阵a*b有意义，否则会panic
// 用户二维数组表示矩阵，矩阵乘法运算如下
// A = [3][2]int64, B = [2][3]int64
// C = AB = [2][2]int64{
// [2]int64{A[0][0]*B[0][0]+A[0][1]*B[1][0]+A[0][2]*B[2][0], A[0][0]*B[0][1]+A[0][1]*B[1][1]+A[0][2]*B[2][1]}
// [2]int64{A[1][0]*B[0][0]+A[0][1]*B[1][0]+A[1][2]*B[2][0], A[1][0]*B[0][1]+A[1][1]*B[1][1]+A[1][2]*B[2][1]}
//}
func MatrixMulti(a, b [][]int64) [][]int64 {
	if len(a) == 0 && len(b) == 0 {
		return nil
	}
	if len(a) == 0 || len(b) == 0 {
		panic("无意义乘法运算")
	}

	aCol := len(a[0])
	bRow := len(b)

	if aCol != bRow {
		panic("无意义乘法运算")
	}

	c := make([][]int64, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = make([]int64, len(b[0]))
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return c
}
