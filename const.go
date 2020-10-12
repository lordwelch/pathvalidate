package pathvalidate

import (
	"errors"
	"unicode"

	"golang.org/x/text/unicode/rangetable"
)

var (
	NTFSReserved = []string{
		"$MFT",
		"$MFTMIRR",
		"$LOGFILE",
		"$VOLUME",
		"$ATTRDEF",
		"$BITMAP",
		"$BOOT",
		"$BADCLUS",
		"$SECURE",
		"$UPCASE",
		"$EXTEND",
		"$QUOTA",
		"$OBJID",
		"$REPARSE",
	} // Only in root directory

	WindowsReserved = []string{
		"CON", "PRN", "AUX", "CLOCK$", "NUL",
		"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9", "COM10",
		"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9", "LPT10",
	}

	DarwinReserved = []string{":"} // Is this needed?
)

var (
	InvalidPath              = rangetable.Merge(unicode.Cc, unicode.Cf, unicode.Z)
	InvalidFilename          = rangetable.Merge(InvalidPath, rangetable.New('/'))
	InvalidWindowsPath       = rangetable.Merge(InvalidPath, rangetable.New(':', '*', '?', '"', '<', '>', '|'))
	InvalidWindowsFilename   = rangetable.Merge(InvalidFilename, InvalidWindowsPath, rangetable.New('\\'))
	DefaultMaxFilenameLength = 255
)

var (
	ErrInvalidChar  = errors.New("pathvalidate: invalid character")
	ErrMaxLength    = errors.New("pathvalidate: max length exceeded")
	ErrMinLength    = errors.New("pathvalidate: min length not met")
	ErrReservedWord = errors.New("pathvalidate: reserved word found")
)
