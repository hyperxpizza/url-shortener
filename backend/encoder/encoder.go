package encoder

import "strings"

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
