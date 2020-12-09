package metset

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// Tmpl is the core object to inspect
type Tmpl struct {
	filePath string
}

// New is the Tmpl builder
func New(inPath string) Tmpl {
	return Tmpl{filePath: inPath}
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

	rxp, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
		return false
	}

	fbytes, err := ioutil.ReadFile(t.filePath)
	if err != nil {
		log.Println(err)
		return false
	}

	slice := strings.Split(string(fbytes), " ")

	for v := range slice {
		if rxp.MatchString(slice[v]) {
			return true
		}
	}

	// log.Println(i, "not found")
	return false
}
