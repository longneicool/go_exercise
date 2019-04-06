package popcount

import (
	"fmt"
	"time"
)

var pc [256]byte

// 在ＰＣ中保存的是0~255中bit 1的个数
// 计算公式容易理解：右移一位　＋　最低位
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountBySingle(x uint64) int {
	start := time.Now()
	numBytes := int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])

	fmt.Printf("single: time eclipses: %gs ", time.Since(start).Seconds())

	return numBytes
}

func PopCountByLoop(x uint64) int {
	start := time.Now()
	var numBytes int
	for i := uint(0); i < 8; i++ {
		numBytes += int(pc[byte(x>>(i*8))])
	}

	fmt.Printf("Loop: time elapsed: %gs ", time.Since(start).Seconds())

	return numBytes
}

func PopCountByShift(x uint64) int {
	start := time.Now()
	var numBytes int
	for i := uint(0); i < 64; i++ {
		if (x & (1 << i)) != 0 {
			numBytes++
		}
	}

	fmt.Printf("Shift: time elapsed: %gs ", time.Since(start).Seconds())

	return numBytes
}

func PopCountFromByteSlice(s []byte) int {
	var numBytes int
	for _, b := range s {
		numBytes += int(pc[b])
	}

	return numBytes
}
