package lpenc

import (
	"github.com/lambrospetrou/goencoding/lpenc"
	"testing"
)

const base62Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func TestNewEncoding(t *testing.T) {
	var enc *lpenc.Encoding = lpenc.NewEncoding(base62Chars)
	if enc == nil {
		t.Errorf("nil Encoding object returned on constructor")
	}
}

func TestEncode(t *testing.T) {
	var enc *lpenc.Encoding = lpenc.NewEncoding(base62Chars)
	encText := enc.Encode(0)
	if encText != "A" {
		t.Error("zero should be encoded as empty string")
	}
	encText = enc.Encode(1)
	if encText != "B" {
		t.Error("1 encoded wrongly")
	}
	encText = enc.Encode(61)
	if encText != "9" {
		t.Error("61 encoded wrongly")
	}
	encText = enc.Encode(64)
	if encText != "BC" {
		t.Error("64 encoded wrongly: ", encText)
	}
}

func TestDecode(t *testing.T) {
	var enc *lpenc.Encoding = lpenc.NewEncoding(base62Chars)
	encText, err := enc.Decode("A")
	if encText != 0 || err != nil {
		t.Error("\"A\" should be decoded to 0")
	}
	encText, err = enc.Decode("B")
	if encText != 1 || err != nil {
		t.Error("\"B\" decoded wrongly instead of 1: ", encText)
	}
	encText, err = enc.Decode("9")
	if encText != 61 || err != nil {
		t.Error("\"9\" decoded wrongly instead of 61: ", encText)
	}
	encText, err = enc.Decode("BC")
	if encText != 64 || err != nil {
		t.Error("\"BC\" decoded wrongly instead of 64: ", encText)
	}
}

func TestReverseString(t *testing.T) {
	var s string = "hello, world"
	rs := lpenc.ReverseString(s)
	if rs != "dlrow ,olleh" {
		t.Error(s, " reversed wrongly: ", rs)
	}
	s = ""
	rs = lpenc.ReverseString(s)
	if rs != "" {
		t.Error(s, " reversed wrongly: ", rs)
	}
}
