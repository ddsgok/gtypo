package strategy_test

import (
	"testing"

	"github.com/ddsgok/gtypo/strategy"
)

func TestHomoglyph(t *testing.T) {
	out, err := strategy.Homoglyph.Generate("zenithar", "")
	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	if len(out) != 90 {
		t.FailNow()
	}
}
