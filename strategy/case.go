package strategy

import (
	"fmt"
	"unicode"

	"github.com/ddsgok/gtypo/helpers"
)

var (
	Case Strategy
)

type caseStrategy struct {
}

// -----------------------------------------------------------------------------

func swapRune(r rune) (n rune, ok bool) {
	ok = true
	if unicode.IsLetter(r) {
		if unicode.IsLower(r) {
			n = unicode.ToUpper(r)
		} else {
			n = unicode.ToLower(r)
		}
	} else {
		ok = false
		n = r
	}

	return
}

func (s *caseStrategy) Generate(domain, tld string) ([]string, error) {
	res := []string{}

	dom := []rune(domain)
	for ws := 0; ws <= len(dom) + 1; ws++{                         // [0, 1, 2, ...]
		for i := 0; i <= (len(dom) - ws); i++ {   // ws = 2 -> [0, 1, 2, 3]
			win := dom[i : i+ws]

			j := 0
			for j < ws {
				c := rune(win[j])
				if rep, ok := swapRune(c); ok {
					win = []rune(fmt.Sprintf("%s%s%s", string(win[:j]), string(rep), string(win[j+1:])))
					fuzzed := fmt.Sprintf("%s%s%s", string(dom[:i]), string(win), string(dom[i+ws:]))
					fuzzed = combineTLD(fuzzed, tld)
					res = append(res, fuzzed)
				}
				j++
			}
		}
	}

	return helpers.Dedup(res), nil
}

func (s *caseStrategy) GetName() string {
	return "case"
}

func init() {
	Case = &caseStrategy{}
}
