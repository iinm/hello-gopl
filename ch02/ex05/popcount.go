package popcount

// pc[i]はiのポピュレーションカウント
var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount xのポピュレーションカウントを返す
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountByLoop(x uint64) int {
	var count byte = 0
	var i uint8
	for i = 0; i < 8; i++ {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}

func PopCountByShift(x uint64) int {
	count := 0
	for i := uint8(0); i < 64; i++ {
		count += int((x >> i) & 1)
	}
	return count
}

func PopCountByClear(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
