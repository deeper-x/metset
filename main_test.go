package metset

import (
	"testing"
)

func TestBasketIsNotMet(t *testing.T) {
	f := New("./assets/file.html")

	if f.BasketVarIsMet([]string{"sed", "anim", "consequat", "boo"}) {
		t.Error("all except boo should be met")
	}
}

func TestBasketIsMet(t *testing.T) {
	f := New("./assets/file.html")

	if !f.BasketVarIsMet([]string{"sed", "anim", "consequat", "culpa"}) {
		t.Error("all vars should be met")
	}
}

func TestContainsAnim(t *testing.T) {
	f := New("./assets/file.html")

	if !f.VarIsMet("anim") {
		t.Error("{{.anim}} not found, even if is an existent variable")
	}
}

func TestContainsSed(t *testing.T) {
	f := New("./assets/file.html")

	if !f.VarIsMet("sed") {
		t.Error("{{.sed}} not found, even if is an existent variable")
	}
}

func TestNotContains(t *testing.T) {
	f := New("./assets/file.html")

	if f.VarIsMet("boo") {
		t.Error("{{.boo}} is unexistent...")
	}
}
