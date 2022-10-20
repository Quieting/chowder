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
