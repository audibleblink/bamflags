package bamflags

import (
	"fmt"
	"math"
	"strings"
)

const (
	off = "0"
	on  = "1"
)

// ParseInt will return bit values of which the Binary Alignment Map is comprised
// https://en.wikipedia.org/wiki/SAM_(file_format)#Bitwise_flags
// Example: ParseInt(514)=> [ 512, 2 ]
func ParseInt(bamInt int64) (opts []int, err error) {
	binary := fmt.Sprintf("%b", bamInt)
	splitBin := strings.Split(binary, "")

	for idx, bit := range splitBin {
		if bit == on {
			sig := len(splitBin) - idx
			opts = append(opts, Significance(sig))
		}
	}
	return
}

// Significance will provide the value at the n-th column of a binary number
// Example: Significance(4) == 8
// | column       | 4 | 3 | 2 | 1 |
// | significance | 8 | 4 | 2 | 1 |
// | binary #     | 1 | 0 | 0 | 1 |
func Significance(column int) int {
	num := float64(column)
	out := math.Exp2(num - 1)
	return int(out)
}
