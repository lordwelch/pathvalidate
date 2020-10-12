package main

import (
	"fmt"
	"os"

	"github.com/lordwelch/pathvalidate"
)

func main() {
	fmt.Println(pathvalidate.ValidateFilepath(os.Args[1]))
	fmt.Println(pathvalidate.SanitizeFilepath(os.Args[1], '_'))
}
