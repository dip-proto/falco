package shared

import (
	"bytes"
	"encoding/base64"
	"errors"
)

var nullByte = []byte{0}

var ErrInvalidBase64 = errors.New("invalid base64 input")

func terminateNullByte(decoded []byte) ([]byte, bool) {
	before, _, found := bytes.Cut(decoded, nullByte)
	return before, found
}

func addPadding(s string) string {
	switch len(s) % 4 {
	case 2:
		return s + "=="
	case 3:
		return s + "="
	default:
		return s
	}
}

func isValidBase64Char(b byte) bool {
	return (b >= 'A' && b <= 'Z') ||
		(b >= 'a' && b <= 'z') ||
		(b >= '0' && b <= '9') ||
		b == '+' || b == '/' || b == '='
}

func isValidBase64UrlChar(b byte) bool {
	return (b >= 'A' && b <= 'Z') ||
		(b >= 'a' && b <= 'z') ||
		(b >= '0' && b <= '9') ||
		b == '-' || b == '_' || b == '='
}

func isValidBase64UrlNoPadChar(b byte) bool {
	return (b >= 'A' && b <= 'Z') ||
		(b >= 'a' && b <= 'z') ||
		(b >= '0' && b <= '9') ||
		b == '-' || b == '_'
}

func validateBase64Std(src string) error {
	for i := 0; i < len(src); i++ {
		if !isValidBase64Char(src[i]) {
			return ErrInvalidBase64
		}
	}
	return nil
}

func validateBase64Url(src string) error {
	for i := 0; i < len(src); i++ {
		if !isValidBase64UrlChar(src[i]) {
			return ErrInvalidBase64
		}
	}
	return nil
}

func validateBase64UrlNoPad(src string) error {
	for i := 0; i < len(src); i++ {
		if !isValidBase64UrlNoPadChar(src[i]) {
			return ErrInvalidBase64
		}
	}
	return nil
}

func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func Base64UrlEncode(src string) string {
	return base64.URLEncoding.EncodeToString([]byte(src))
}

func Base64UrlEncodeNoPad(src string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(src))
}

type Base64DecodeResult struct {
	Value              string
	HadEmbeddedNulByte bool
}

func Base64Decode(src string) (Base64DecodeResult, error) {
	if src == "" {
		return Base64DecodeResult{Value: ""}, nil
	}

	if err := validateBase64Std(src); err != nil {
		return Base64DecodeResult{}, err
	}

	dec, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return Base64DecodeResult{}, ErrInvalidBase64
	}

	result, hadNull := terminateNullByte(dec)
	return Base64DecodeResult{
		Value:              string(result),
		HadEmbeddedNulByte: hadNull,
	}, nil
}

func Base64UrlDecode(src string) (Base64DecodeResult, error) {
	if src == "" {
		return Base64DecodeResult{Value: ""}, nil
	}

	if err := validateBase64Url(src); err != nil {
		return Base64DecodeResult{}, err
	}

	dec, err := base64.URLEncoding.DecodeString(src)
	if err != nil {
		return Base64DecodeResult{}, ErrInvalidBase64
	}

	result, hadNull := terminateNullByte(dec)
	return Base64DecodeResult{
		Value:              string(result),
		HadEmbeddedNulByte: hadNull,
	}, nil
}

func Base64UrlDecodeNoPad(src string) (Base64DecodeResult, error) {
	if src == "" {
		return Base64DecodeResult{Value: ""}, nil
	}

	if err := validateBase64UrlNoPad(src); err != nil {
		return Base64DecodeResult{}, err
	}

	padded := addPadding(src)

	dec, err := base64.RawURLEncoding.DecodeString(padded)
	if err != nil {
		dec, err = base64.RawURLEncoding.DecodeString(src)
		if err != nil {
			return Base64DecodeResult{}, ErrInvalidBase64
		}
	}

	result, hadNull := terminateNullByte(dec)
	return Base64DecodeResult{
		Value:              string(result),
		HadEmbeddedNulByte: hadNull,
	}, nil
}
