package encoder

import (
	"fmt"
	"math"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint64(len(alphabet))
)

func Encode(n uint64) string {
	var builder strings.Builder
	builder.Grow(11)

	for ; n > 0; n = n / length {
		builder.WriteByte(alphabet[n%length])
	}

	return builder.String()
}

func Decode(encoded string) (uint64, error) {
	var n uint64

	for i, v := range encoded {
		pos := strings.IndexRune(alphabet, v)
		if pos == -1 {
			return uint64(pos), fmt.Errorf("invalid character: %s\n", string(v))
		}

		n += uint64(pos) * uint64(math.Pow(float64(length), float64(i)))
	}

	return n, nil
}
