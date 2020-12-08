package metset

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// Tmpl is the core object to inspect
type Tmpl struct {
	filePath *os.File
}

// Open is the tmpl builder
func Open(path string) Tmpl {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return Tmpl{filePath: f}
}

// Close close the file handler
func (t Tmpl) Close() {
	t.filePath.Close()
}

// Contains says if tmpl variable is found
func (t Tmpl) Contains(i string) bool {
	// searching for {{ .var }} or {{.var}} or {{.var }} or {{ .var}}
	pattern := fmt.Sprintf(`\{\{\ {0,1}\.%v\ {0,1}\}\}`, i)
	scanner := bufio.NewScanner(t.filePath)

	rxp, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
		return false
	}

	// scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		if rxp.MatchString(scanner.Text()) {
			return true
		}
	}

	return false
}
