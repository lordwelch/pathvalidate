package pathvalidate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	DefaultFilenameSanitizer = FilenameSanitizer{}
	DefaultFilepathSanitizer = FilepathSanitizer{}
)

type FilepathSanitizer struct {
	FilenameSanitizer
}

func (fps FilepathSanitizer) Sanitize(path string, replacement rune) (string, error) {
	var (
		err error
	)
	cleaned := filepath.Clean(path)
	split := strings.Split(cleaned, string(os.PathSeparator))
	splitS := make([]string, 0, len(split))
	for _, name := range split {
		name, err = fps.FilenameSanitizer.Sanitize(name, replacement)
		if err != nil {
			return path, err
		}
		splitS = append(splitS, name)
	}
	return filepath.Join(splitS...), nil
}

func (fps FilepathSanitizer) Validate(path string) error {
	cleaned := filepath.Clean(path)
	split := strings.Split(cleaned, string(os.PathSeparator))
	for _, name := range split {
		if err := fps.FilenameSanitizer.Validate(name); err != nil {
			return err
		}
	}
	return nil
}

type FilenameSanitizer struct {
	BaseFile
}

func (f FilenameSanitizer) Sanitize(path string, replacement rune) (string, error) {
	var (
		err error
	)
	if f.BaseFile.MinLength == 0 {
		f.BaseFile = DefaultBaseFile
	}
	replace := func(r rune) rune {
		if unicode.Is(InvalidFilename, r) && r != ' ' {
			return replacement
		}
		return r
	}
	sanitized := strings.Map(replace, path)
	sanitized = f.UpdateReservedKeywords(sanitized, "_")
	sanitized = strings.TrimSpace(sanitized)
	err = f.Validate(sanitized)
	if err != nil {
		return path, fmt.Errorf("could not validate sanitized filename: %w", err)
	}
	return sanitized, nil
}

func (f FilenameSanitizer) Validate(path string) error {
	if f.BaseFile.MinLength == 0 {
		f.BaseFile = DefaultBaseFile
	}
	nameLen := utf8.RuneCountInString(path)
	cleaned := filepath.Clean(path)

	if nameLen > f.MaxLength {
		return fmt.Errorf("%w: wanted <= %d, got = %d", ErrMaxLength, f.MaxLength, nameLen)
	}

	if nameLen < f.MinLength {
		return fmt.Errorf("%w: wanted >= %d, got = %d", ErrMinLength, f.MinLength, nameLen)
	}

	err := f.validateReservedKeywords(cleaned)
	if err != nil {
		return err
	}

	validate := func(r rune) bool {
		return unicode.Is(InvalidFilename, r) && r != ' '
	}
	if n := strings.IndexFunc(cleaned, validate); n != -1 {
		r, _ := utf8.DecodeRuneInString(cleaned[n:])
		return fmt.Errorf("%w: '%s' (%#x)", ErrInvalidChar, string(r), r)
	}
	if cleaned[0] == ' ' {
		return fmt.Errorf("%w: space at beginning of string", ErrInvalidChar)
	}
	if cleaned[len(cleaned)-1] == ' ' {
		return fmt.Errorf("%w: space at end of string", ErrInvalidChar)
	}
	return nil
}

func SanitizeFilename(path string, replacement rune) (string, error) {
	return DefaultFilenameSanitizer.Sanitize(path, replacement)
}

func ValidateFilename(path string) error {
	return DefaultFilenameSanitizer.Validate(path)
}

func SanitizeFilepath(path string, replacement rune) (string, error) {
	return DefaultFilepathSanitizer.Sanitize(path, replacement)
}

func ValidateFilepath(path string) error {
	return DefaultFilepathSanitizer.Validate(path)
}
