package match

// Matcher 游戏匹配算法的一部分，将参与匹配的人按照目标人数匹配成组
type Matcher interface {
	Number() int // 包含人数
}

// Match 多人匹配成 n 一组，目前仅支持6人匹配
func Match(data []Matcher, n int) [][]Matcher {
	if n != 6 {
		return nil
	}
	res := make([][]Matcher, 0)

	// 按人数将人数分组
	var groups = make([][]Matcher, n-1)
	for _, val := range data {
		num := val.Number()
		if num > n {
			continue
		}
		if num == n {
			res = append(res, []Matcher{val})
			continue
		}
		groups[num-1] = append(groups[num-1], val)
	}

	// 先两两成组匹配成功
	gLen := len(groups)
	for i := 0; i <= (gLen-1)/2; i++ {
		minLen := min(len(groups[i]), len(groups[gLen-1-i]))
		if i == (gLen-1)/2 {
			minLen /= 2
		}

		// 如果切片长度为奇数，将对中间切片元素进行两次切割
		max := groups[i][:minLen]
		groups[i] = groups[i][minLen:]

		min := groups[gLen-1-i][:minLen]
		groups[gLen-1-i] = groups[gLen-1-i][minLen:]

		for j := 0; j < minLen; j++ {
			res = append(res, []Matcher{max[j], min[j]})
		}
	}

	// 将数据按从大到小排序
	data = data[:0]
	for i := n - 2; i >= 0; i-- {
		data = append(data, groups[i]...)
	}

	head, tail := 0, len(data)-1 // 头尾指针
	isHead := true
	item := make([]Matcher, 1) // 头指针元素放在第一位
	for j := 0; j < len(data); j++ {
		if isHead {
			item[0] = data[head]
			head++
		} else {
			item = append(item, data[tail])
			tail--
		}

		switch total := sum(item); {
		case total == n:
			res = append(res, append([]Matcher{}, item...))
			item = item[:1]

			isHead = true
		case total < n:
			isHead = false
		default:
			for i, val := range item {
				if val.Number() != total-n {
					continue
				}
				tmp := make([]Matcher, 0, len(item)-1)
				tmp = append(tmp, item[:i]...)
				tmp = append(tmp, item[i+1:]...)
				res = append(res, tmp)

				item = item[:1]
			}

			isHead = true
		}
	}

	return res
}

func sum(v []Matcher) int {
	s := 0
	for _, val := range v {
		s += val.Number()
	}

	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
