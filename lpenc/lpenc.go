package lpenc

import (
	"bytes"
	"errors"
	"fmt"
	//"strconv"
)

// 62 valid characters a-z0-9A-Z
const base62Chars = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 64 valid characters A-Za-z0-9-_
const base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

type Encoding struct {
	encodeLen   int
	encodeChars string
	decodeMap   [256]byte
}

func NewEncoding(chars string) *Encoding {
	e := new(Encoding)
	e.encodeChars = chars
	e.encodeLen = len(chars)
	map_sz := len(e.decodeMap)
	// initialize decode map - this will be used as an error check value
	for i := 0; i < map_sz; i++ {
		e.decodeMap[i] = 0xFF
	}
	// create the correct decode map
	for i, c := range chars {
		e.decodeMap[c] = byte(i)
	}
	return e
}

func (enc *Encoding) Encode(n uint64) string {
	var buffer bytes.Buffer
	for n > 0 {
		mod := n % uint64(enc.encodeLen)
		buffer.WriteByte(enc.encodeChars[mod])
		n /= uint64(enc.encodeLen)
	}
	return buffer.String()
}

func (enc *Encoding) Decode(s string) (uint64, error) {
	var n uint64 = 0
	var b byte
	for _, c := range []byte(s) {
		b = enc.decodeMap[c]
		if b == 0xFF {
			return 0, errors.New("Contains invalid char")
		}
		n *= uint64(enc.encodeLen)
		n += uint64(enc.decodeMap[c])
	}
	return n, nil
}

// 62 valid characters a-z0-9A-Z
var Base62Encoding = NewEncoding(base62Chars)

// 64 valid characters a-z0-9A-Z
var Base64Encoding = NewEncoding(base64Chars)
