package pathvalidate_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/lordwelch/pathvalidate"
)

var tests = []struct {
	path, sanitized string
	err             error
}{
	{"hello\t", "hello_", pathvalidate.ErrInvalidChar},
	{"hello\r", "hello_", pathvalidate.ErrInvalidChar},
	{"hello\n", "hello_", pathvalidate.ErrInvalidChar},
	{"hello ", "hello", pathvalidate.ErrInvalidChar},
	{"hello/world", "hello_world", pathvalidate.ErrInvalidChar},
	{"nul", "nul_", pathvalidate.ErrReservedWord},
	{"nul.test", "nul_.test", pathvalidate.ErrReservedWord},
	{"hello" + strings.Repeat(" ", 4090) + "world", "hello" + strings.Repeat(" ", 4090) + "world", pathvalidate.ErrMaxLength},
	{"", "", pathvalidate.ErrMinLength},
	{"hello world", "hello world", nil},
}

func TestValidate(t *testing.T) {
	pathvalidate.DefaultBaseFile.ReservedKeywords = pathvalidate.WindowsReserved
	for _, test := range tests {
		if err := pathvalidate.ValidateFilename(test.path); !errors.Is(err, test.err) {
			t.Errorf("got %v, want %v", err, test.err)
		}
	}
}

func TestSanitize(t *testing.T) {

	for _, test := range tests {
		// Skips length tests as there is no way to intelligently sanitize them
		if errors.Is(test.err, pathvalidate.ErrMaxLength) || errors.Is(test.err, pathvalidate.ErrMinLength) {
			continue
		}
		if got, err := pathvalidate.SanitizeFilename(test.path, '_'); err != nil || got != test.sanitized {
			t.Errorf("got value: %v; error: %v, want value: %v; error: %v", got, err, test.sanitized, nil)
		}
	}
}
