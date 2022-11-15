package try

import (
	"encoding/json"
)

func openDoor(d []int64) []int64 {
	keys := make(map[int64]struct{}, len(d)) // 已找到的钥匙
	nextOpenDoor := int64(1)                 // 最近打开的门
	for i, key := range d {
		keys[key] = struct{}{}
		if key != nextOpenDoor {
			continue
		}
		d[nextOpenDoor-1] = int64(i + 1)
		for {
			nextOpenDoor++
			_, ok := keys[nextOpenDoor]
			if !ok {
				break
			}
			d[nextOpenDoor-1] = int64(i + 1)
		}

	}
	return d
}

func equal(s string) string {
	return ""
}

func parse(s string) (action []rune, nums []int64) {
	action = make([]rune, 0)
	nums = make([]int64, 0)
	n := int64(0)
	for _, l := range s {
		switch l {
		case '+', '*':
			action = append(action, l)
			nums = append(nums, n)
			n = 0
		case '=':
		default:
			n = n*10 + int64(l-'0')
		}
	}

	return
}

type GiveGiftArg struct{}
type SettleArg struct{}
type Register interface {
	ConsumeGiveGift(key string, fn func(arg GiveGiftArg))
	ConsumeSettle(key string, fn func(arg SettleArg))

	SendGiveGift(key string, arg GiveGiftArg)
}

// 注册 kafka
func init() {
}

// RegisterUser 对接kafka注册方法
func RegisterUser(body string) {
	msg := struct {
		key string
		arg interface{}
	}{}

	json.Unmarshal([]byte(body), &msg)

	switch msg.key {
	case "key1":

	}
}

type RoomInfo struct {
	Id     int64
	People int64
}

func Match(data []RoomInfo) [][]RoomInfo {
	res := make([][]RoomInfo, 0, len(data))

	// 按人数将房间分组
	var groups = [6][]RoomInfo{}
	for _, val := range data {
		groups[val.People-1] = append(groups[val.People-1], val)
	}

	// 现将33，42， 51 房间匹配成组
	for i, j := 0, len(groups[2])-1; i < j; {
		res = append(res, []RoomInfo{
			groups[2][i],
			groups[2][j],
		})
		i++
		j--
	}

	// 将剩余房间按人数排序
	data = data[:0]
	for i := 5; i >= 0; i-- {
		list := groups[i]
		if i == 2 {
			if len(list)%2 == 1 {
				data = append(data, list[len(list)/2])
			}
			continue
		}
		for _, val := range list {
			data = append(data, val)
		}
	}

	// 将剩余房间匹配
	i := 0
	sort := 1 // 正数处于队头，负数处于队尾
	item := make([]RoomInfo, 0, 6)
	for j := 0; j < len(data); j++ {
		item = append(item, data[i])
		total := int64(0)
		for _, val := range item {
			total += val.People
		}
		switch {
		case total == 6:
			res = append(res, item)
			item = []RoomInfo{}

			if sort > 0 {
				i++
			} else {
				i = len(data) - i // 将i置于队头
				sort = -sort
			}
		case total < 6:
			if sort > 0 {
				i = len(data) - i - 1 // 将i置于队尾
				sort = -sort
			} else {
				i--
			}
		case total == 7:
			for i, val := range item {
				if val.People == 1 {
					v := append([]RoomInfo{}, item[:i]...)
					v = append(v, item[i+1:]...)
					res = append(res, v)

					item = []RoomInfo{}
					break
				}
			}
			if sort > 0 {
				i++
			} else {
				i = len(data) - i // 将i置于队头
				sort = -sort
			}
		default:
			item = append([]RoomInfo{}, item[1:]...)
			if sort > 0 {
				i = len(data) - 1 - i // 将i置于队尾
				sort = -sort
			} else {
				i--
			}
		}
	}

	return res
}
