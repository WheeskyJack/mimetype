package matchers

import "bytes"

func Pdf(in []byte) bool {
	return bytes.Equal(in[:4], []byte{0x25, 0x50, 0x44, 0x46})
}
