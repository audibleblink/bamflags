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

	// Start from the end (least significant bits), so the
	// resulting slice is in ascending bit order
	for idx := len(splitBin) - 1; idx >= 0; idx-- {
		bit := splitBin[idx]
		if bit == on {
			column := len(splitBin) - idx
			opts = append(opts, Significance(column))
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
	out := math.Exp2(num - 1) // 2 ^ column
	return int(out)
}

// Contains will tell the caller whether or not a certain flag
// within the provided Binary Access Map is set
func Contains(binMap, bitValue int64) (isSet bool) {
	return binMap&bitValue != 0
}
