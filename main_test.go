package metset

import (
	"testing"
)

func TestContainsAnim(t *testing.T) {
	f := Open("./assets/file.txt")
	defer f.Close()

	if !f.Contains("anim") {
		t.Error("{{.anim}} not found, even if is an existent variable")
	}
}

func TestContainsSed(t *testing.T) {
	f := Open("./assets/file.txt")
	defer f.Close()

	if !f.Contains("sed") {
		t.Error("{{.sed}} not found, even if is an existent variable")
	}
}

func TestNotContains(t *testing.T) {
	f := Open("./assets/file.txt")
	defer f.Close()

	if f.Contains("boo") {
		t.Error("{{.boo}} is unexistent...")
	}
}
