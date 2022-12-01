package try

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
