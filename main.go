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

// BasketVarIsMet for a given input slice, check for all vars
func (t Tmpl) BasketVarIsMet(vars []string) bool {
	for k := range vars {
		if !t.VarIsMet(vars[k]) {
			return false
		}
	}

	return true
}

// VarIsMet says if tmpl's variable is found
func (t Tmpl) VarIsMet(i string) bool {
	// searching for {{ .var }} or {{.var}} or {{.var }} or {{ .var}}
	pattern := fmt.Sprintf(`\{\{\s{0,1}\.%v\s{0,1}\}\}`, i)
	scanner := bufio.NewScanner(t.filePath)

	rxp, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
		return false
	}

	// scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		if rxp.MatchString(scanner.Text()) {
			log.Println(scanner.Text(), "found")
			return true
		}
	}

	return false
}
