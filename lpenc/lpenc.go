package lpenc

import (
	"bytes"
	"errors"
	"unicode/utf8"
)

// 62 valid characters A-Za-z0-9
const base62Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

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

func ReverseString(s string) string {
	cs := make([]rune, utf8.RuneCountInString(s))
	i := len(cs)
	for _, c := range s {
		i--
		cs[i] = c
	}
	return string(cs)
}

// returns the String representation of the uint64 passed in.
// empty string ("") is returned if the value is 0
func (enc *Encoding) Encode(n uint64) string {
	var buffer bytes.Buffer
	for n > 0 {
		mod := n % uint64(enc.encodeLen)
		buffer.WriteByte(enc.encodeChars[mod])
		n /= uint64(enc.encodeLen)
	}
	return ReverseString(buffer.String())
}

// tries to decode the string passed in into a uint64.
// if the string contains a character not defined in the alphabet of the
// encoding in use then an error is returned, uint64 value is undefined what will have
func (enc *Encoding) Decode(s string) (uint64, error) {
	var n uint64 = 0
	var b byte
	for _, c := range []byte(ReverseString(s)) {
		b = enc.decodeMap[c]
		if b == 0xFF {
			return 0, errors.New("Contains invalid char")
		}
		n *= uint64(enc.encodeLen)
		n += uint64(enc.decodeMap[c])
	}
	return n, nil
}

// 62 valid characters A-Za-z0-9
var Base62Encoding = NewEncoding(base62Chars)

// 64 valid characters A-Za-z0-9-_
var Base64Encoding = NewEncoding(base64Chars)
