package smslen

import (
	"regexp"
	"unicode/utf8"
)

type encoding string

const (
	gsm7Bit    encoding = "GSM_7BIT"
	gsm7BitExt encoding = "GSM_7BIT_EXT"
	utf16      encoding = "UTF_16"
)

var (
	gsm7BitRegex         = regexp.MustCompile(`^[@£$¥èéùìòÇ\nØø\rÅåΔ_ΦΓΛΩΠΨΣΘΞÆæßÉ !"#¤%&'()*+,\-./0123456789:;<=>?¡ABCDEFGHIJKLMNOPQRSTUVWXYZÄÖÑÜ§¿abcdefghijklmnopqrstuvwxyzäöñüà]*$`)
	gsm7BitExtRegex      = regexp.MustCompile(`^[@£$¥èéùìòÇ\nØø\rÅåΔ_ΦΓΛΩΠΨΣΘΞÆæßÉ !"#¤%&'()*+,\-./0123456789:;<=>?¡ABCDEFGHIJKLMNOPQRSTUVWXYZÄÖÑÜ§¿abcdefghijklmnopqrstuvwxyzäöñüà^{}\\[~\]|€]*$`)
	gsm7BitExtCharsRegex = regexp.MustCompile(`[\^{}\\[~\]|€]`)

	partLength = map[encoding]int{
		gsm7Bit:    160,
		gsm7BitExt: 160,
		utf16:      70,
	}

	multiPartLength = map[encoding]int{
		gsm7Bit:    153,
		gsm7BitExt: 153,
		utf16:      67,
	}
)

type CountResult struct {
	Encoding    string
	Chars       int
	CharsInPart int
	Parts       int
}

func detectEncoding(text string) encoding {
	if gsm7BitRegex.MatchString(text) {
		return gsm7Bit
	}

	if gsm7BitExtRegex.MatchString(text) {
		return gsm7BitExt
	}

	return utf16
}

func countLength(s string, e encoding) int {
	length := utf8.RuneCountInString(s)

	if e == gsm7BitExt {
		length += len(gsm7BitExtCharsRegex.FindAllString(s, -1))
	}

	return length
}

func countCharsInPart(length int, encoding encoding) int {
	if length > partLength[encoding] {
		return multiPartLength[encoding]
	}

	return partLength[encoding]
}

func countParts(length int, charsInPart int) int {
	return (length + charsInPart - 1) / charsInPart
}

func Count(text string) CountResult {
	encoding := detectEncoding(text)
	length := countLength(text, encoding)
	charsInPart := countCharsInPart(length, encoding)
	parts := countParts(length, charsInPart)

	return CountResult{
		Encoding:    string(encoding),
		Chars:       length,
		CharsInPart: charsInPart,
		Parts:       parts,
	}
}
