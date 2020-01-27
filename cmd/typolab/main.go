package main

import (
	"fmt"
	"os"

	"github.com/ddsgok/gtypo"
	"github.com/ddsgok/gtypo/strategy"
)

func main() {
	args := os.Args

	all := []strategy.Strategy{
		strategy.Case,
		strategy.Addition,
		strategy.Homoglyph,
		strategy.Omission,
	}
	
	results, err := gtypo.Fuzz(args[1], all...)
	if err != nil {
		fmt.Println(err)
	}
	
	for _, r := range results {
		fmt.Println(append([]string{r.Domain}, r.Permutations...))
	}

}
