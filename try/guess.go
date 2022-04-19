package try

import "strconv"

func guess(opt []string, nums []int, want int) string {

	return ""

}

// merge 合并 a, b
// 例如: a = 10, b = 1 合并后得到 101
func merge(a, b int) int {
	m, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return m
}

// arrangement 返回组合得到的种数
func arrangement(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{nums[0]})
	for _, num := range nums[1:] {
		ans = arrangement1(ans, num)
	}

	return ans
}

func arrangement1(ans [][]int, num int) [][]int {
	res := make([][]int, 0)
	for _, val := range ans {
		for i := 0; i < len(val)+1; i++ {
			r := make([]int, 0, len(val)+1)
			if i == 0 {
				r = append(append(r, num), val...)
			} else if i == len(val) {
				r = append(append(r, val...), num)
			} else {
				r = append(append(append(r, val[:i]...), num), val[i:]...)
			}

			res = append(res, r)
		}
	}
	return res
}
