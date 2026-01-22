package shared

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBase64Encode(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "Καλώς ορίσατε",
			expect: "zprOsc67z47PgiDOv8+Bzq/Pg86xz4TOtQ==",
		},
	}

	for _, tt := range tests {
		v := Base64Encode(tt.input)
		if diff := cmp.Diff(tt.expect, v); diff != "" {
			t.Errorf("Return value unmach, diff=%s", diff)
		}
	}
}

func TestBase64UrlEncode(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "Καλώς ορίσατε",
			expect: "zprOsc67z47PgiDOv8-Bzq_Pg86xz4TOtQ==",
		},
	}

	for _, tt := range tests {
		v := Base64UrlEncode(tt.input)
		if diff := cmp.Diff(tt.expect, v); diff != "" {
			t.Errorf("Return value unmach, diff=%s", diff)
		}
	}
}

func TestBase64UrlEncodeNoPad(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "Καλώς ορίσατε",
			expect: "zprOsc67z47PgiDOv8-Bzq_Pg86xz4TOtQ",
		},
	}

	for _, tt := range tests {
		v := Base64UrlEncodeNoPad(tt.input)
		if diff := cmp.Diff(tt.expect, v); diff != "" {
			t.Errorf("Return value unmach, diff=%s", diff)
		}
	}
}

func TestBase64Decode(t *testing.T) {
	t.Run("valid inputs", func(t *testing.T) {
		tests := []struct {
			input  string
			expect string
		}{
			{
				input:  "zprOsc67z47PgiDOv8+Bzq/Pg86xz4TOtQ==",
				expect: "Καλώς ορίσατε",
			},
			{
				input:  "c29tZSBkYXRhIHdpdGggACBhbmQg77u/",
				expect: "some data with ",
			},
			{
				input:  "YWJjZB==",
				expect: "abcd",
			},
			{
				input:  "aGVsbG8=",
				expect: "hello",
			},
		}

		for _, tt := range tests {
			result, err := Base64Decode(tt.input)
			if err != nil {
				t.Errorf("Unexpected error for input %q: %v", tt.input, err)
			}
			if diff := cmp.Diff(tt.expect, result.Value); diff != "" {
				t.Errorf("Return value unmatch, diff=%s", diff)
			}
		}
	})

	t.Run("invalid inputs return error", func(t *testing.T) {
		tests := []string{
			"QU&|*#()JDRA==",
			"QU&==|*#()JDRA==",
			"QU&=|*#()JDRA==",
			"aGVsbG8=0",
		}

		for _, input := range tests {
			_, err := Base64Decode(input)
			if err == nil {
				t.Errorf("Expected error for invalid input %q", input)
			}
		}
	})
}

func TestBase64UrlDecode(t *testing.T) {
	t.Run("valid inputs", func(t *testing.T) {
		tests := []struct {
			input  string
			expect string
		}{
			{
				input:  "zprOsc67z47PgiDOv8-Bzq_Pg86xz4TOtQ==",
				expect: "Καλώς ορίσατε",
			},
			{
				input:  "c29tZSBkYXRhIHdpdGggACBhbmQg77u_",
				expect: "some data with ",
			},
			{
				input:  "YWJjZB==",
				expect: "abcd",
			},
			{
				input:  "aGVsbG8=",
				expect: "hello",
			},
		}

		for _, tt := range tests {
			result, err := Base64UrlDecode(tt.input)
			if err != nil {
				t.Errorf("Unexpected error for input %q: %v", tt.input, err)
			}
			if diff := cmp.Diff(tt.expect, result.Value); diff != "" {
				t.Errorf("Return value unmatch, diff=%s", diff)
			}
		}
	})

	t.Run("invalid inputs return error", func(t *testing.T) {
		tests := []string{
			"QU&|*#()JDRA==",
			"QU&==|*#()JDRA==",
			"QU&=|*#()JDRA==",
			"aGVsbG8=0",
		}

		for _, input := range tests {
			_, err := Base64UrlDecode(input)
			if err == nil {
				t.Errorf("Expected error for invalid input %q", input)
			}
		}
	})
}

func TestBase64UrlDecodeNoPad(t *testing.T) {
	t.Run("valid inputs", func(t *testing.T) {
		tests := []struct {
			input  string
			expect string
		}{
			{
				input:  "zprOsc67z47PgiDOv8-Bzq_Pg86xz4TOtQ",
				expect: "Καλώς ορίσατε",
			},
			{
				input:  "c29tZSBkYXRhIHdpdGggACBhbmQg77u_",
				expect: "some data with ",
			},
			{
				input:  "YWJjZA",
				expect: "abcd",
			},
			{
				input:  "aGVsbG8",
				expect: "hello",
			},
		}

		for _, tt := range tests {
			result, err := Base64UrlDecodeNoPad(tt.input)
			if err != nil {
				t.Errorf("Unexpected error for input %q: %v", tt.input, err)
			}
			if diff := cmp.Diff(tt.expect, result.Value); diff != "" {
				t.Errorf("Return value unmatch, diff=%s", diff)
			}
		}
	})

	t.Run("invalid inputs return error", func(t *testing.T) {
		tests := []string{
			"QU&|*#()JDRA",
			"invalid!!chars",
		}

		for _, input := range tests {
			_, err := Base64UrlDecodeNoPad(input)
			if err == nil {
				t.Errorf("Expected error for invalid input %q", input)
			}
		}
	})
}
