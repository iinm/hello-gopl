package main

// pc[i]はiのポピュレーションカウント
var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// 異なるBitの数を数える
func BitDiffCount(c1, c2 *[32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		a := c1[i] ^ c2[i]
		count += int(pc[a])
	}
	return count
}
