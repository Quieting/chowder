package try

func print() uint64 {
	return 4 << (^uintptr(0) >> 63)
}
