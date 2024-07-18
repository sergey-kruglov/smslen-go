package smslen

import (
	"reflect"
	"testing"
)

func TestCountGsm7SinglePart(t *testing.T) {
	text := "Test message. Test message. Test message. Test message. Test message. Test message. Test message. Test message. Test message. Test message. Test message. Test m"
	want := CountResult{
		Encoding:    string(gsm7Bit),
		Chars:       160,
		CharsInPart: partLength[gsm7Bit],
		Parts:       1,
	}
	res := Count(text)
	if !reflect.DeepEqual(res, want) {
		t.Fatalf(`Count(%s) = %v, want match for %v`, text, res, want)
	}
}

func TestCountGsm7MultiPart(t *testing.T) {
	text := "Test message. Test message. Test message. Test message. Test message. Test message. Test message. Test message. Test message. Test message. Test message. Test message."
	want := CountResult{
		Encoding:    string(gsm7Bit),
		Chars:       167,
		CharsInPart: multiPartLength[gsm7Bit],
		Parts:       2,
	}
	res := Count(text)
	if !reflect.DeepEqual(res, want) {
		t.Fatalf(`Count(%s) = %v, want match for %v`, text, res, want)
	}
}

func TestCountGsm7ExtSinglePart(t *testing.T) {
	text := "^Test message^^Test message^^Test message^^Test message^^Test message^^Test message^^Test message^^Test message^^Test message^^Test message^"
	want := CountResult{
		Encoding:    string(gsm7BitExt),
		Chars:       160,
		CharsInPart: partLength[gsm7BitExt],
		Parts:       1,
	}
	res := Count(text)
	if !reflect.DeepEqual(res, want) {
		t.Fatalf(`Count(%s) = %v, want match for %v`, text, res, want)
	}
}

func TestCountGsm7ExtMultiPart(t *testing.T) {
	text := "^Test message^^Test message^^Test message^^Test message^^Test message^^Test message^^Test message^^Test message^^Test message^^Test message^^Test message^"
	want := CountResult{
		Encoding:    string(gsm7BitExt),
		Chars:       176,
		CharsInPart: multiPartLength[gsm7BitExt],
		Parts:       2,
	}
	res := Count(text)
	if !reflect.DeepEqual(res, want) {
		t.Fatalf(`Count(%s) = %v, want match for %v`, text, res, want)
	}
}

func TestCountUtf16(t *testing.T) {
	text := "±Test message±±Test message±±Test message±±Test message±±Test message±"
	want := CountResult{
		Encoding:    string(utf16),
		Chars:       70,
		CharsInPart: partLength[utf16],
		Parts:       1,
	}
	res := Count(text)
	if !reflect.DeepEqual(res, want) {
		t.Fatalf(`Count(%s) = %v, want match for %v`, text, res, want)
	}
}

func TestCountUtf16MultiPart(t *testing.T) {
	text := "±Test message±±Test message±±Test message±±Test message±±Test message±±Test message±"
	want := CountResult{
		Encoding:    string(utf16),
		Chars:       84,
		CharsInPart: multiPartLength[utf16],
		Parts:       2,
	}
	res := Count(text)
	if !reflect.DeepEqual(res, want) {
		t.Fatalf(`Count(%s) = %v, want match for %v`, text, res, want)
	}
}
