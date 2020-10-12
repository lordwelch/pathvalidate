package pathvalidate

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

type BaseFile struct {
	ReservedKeywords []string
	MinLength        int
	MaxLength        int
}

var DefaultBaseFile = BaseFile{
	MaxLength:        getDefaultMaxLength(runtime.GOOS),
	ReservedKeywords: getDefaultKeywords(runtime.GOOS),
	MinLength:        1,
}

func getDefaultKeywords(platform string) []string {
	switch platform {
	case "windows":
		return append(WindowsReserved, NTFSReserved...)
	case "darwin":
		return DarwinReserved
	default:
		return nil
	}
}

func getDefaultMaxLength(platform string) int {
	switch platform {
	case "linux":
		return 4096
	case "windows":
		return 260
	case "darwin":
		return 1024
	default:
		return DefaultMaxFilenameLength
	}
}

func (bf BaseFile) IsReservedKeyword(name string) bool {
	sort.Strings(bf.ReservedKeywords)
	index := sort.SearchStrings(bf.ReservedKeywords, strings.ToUpper(name))
	return index < len(bf.ReservedKeywords) && bf.ReservedKeywords[index] == strings.ToUpper(name)
}

func (bf BaseFile) UpdateReservedKeywords(name, suffix string) string {
	ext := filepath.Ext(name)
	rootName := extractRootName(name)
	if bf.IsReservedKeyword(strings.ToUpper(rootName)) {
		return rootName + suffix + ext
	}

	return name
}

func (bf BaseFile) validateReservedKeywords(name string) error {
	rootName := extractRootName(name)
	if bf.IsReservedKeyword(strings.ToUpper(rootName)) {
		return fmt.Errorf("%w: %s", ErrReservedWord, rootName)
	}

	return nil
}

func extractRootName(path string) string {
	base := filepath.Base(filepath.Clean(path))
	return strings.TrimSuffix(base, filepath.Ext(base))
}
