package strategy_test

import (
	"fmt"
	"testing"

	"github.com/ddsgok/gtypo/strategy"
)

func TestCase(t *testing.T) {
	out, err := strategy.Case.Generate("start", "")
	fmt.Println(out)

	if err != nil {
		t.Fail()
		t.Fatal("Error should not occurs !", err)
	}

	if len(out) == 0 {
		t.FailNow()
	}

	// if len(out) != 90 {
	// 	t.FailNow()
	// }
}
