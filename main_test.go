package metset

import (
	"testing"
)

func TestBasketIsMet(t *testing.T) {
	f := Open("./assets/file.html")
	defer f.Close()

	if !f.BasketVarIsMet([]string{"sed", "anim", "consequat"}) {
		t.Error("basket not met")
	}
}

func TestContainsAnim(t *testing.T) {
	f := Open("./assets/file.html")
	defer f.Close()

	if !f.VarIsMet("anim") {
		t.Error("{{.anim}} not found, even if is an existent variable")
	}
}

func TestContainsSed(t *testing.T) {
	f := Open("./assets/file.html")
	defer f.Close()

	if !f.VarIsMet("sed") {
		t.Error("{{.sed}} not found, even if is an existent variable")
	}
}

func TestNotContains(t *testing.T) {
	f := Open("./assets/file.html")
	defer f.Close()

	if f.VarIsMet("boo") {
		t.Error("{{.boo}} is unexistent...")
	}
}
